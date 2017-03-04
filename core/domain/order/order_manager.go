/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-05 17:49
 * description :
 * history :
 */

package order

import (
	"errors"
	"go2o/core/domain/interface/cart"
	"go2o/core/domain/interface/delivery"
	"go2o/core/domain/interface/enum"
	"go2o/core/domain/interface/express"
	"go2o/core/domain/interface/item"
	"go2o/core/domain/interface/member"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/order"
	"go2o/core/domain/interface/payment"
	"go2o/core/domain/interface/product"
	"go2o/core/domain/interface/promotion"
	"go2o/core/domain/interface/shipment"
	"go2o/core/domain/interface/valueobject"
	"go2o/core/infrastructure/domain"
	"time"
)

var _ order.IOrderManager = new(orderManagerImpl)

type orderManagerImpl struct {
	repo         order.IOrderRepo
	productRepo  product.IProductRepo
	cartRepo     cart.ICartRepo
	goodsRepo    item.IGoodsItemRepo
	promRepo     promotion.IPromotionRepo
	memberRepo   member.IMemberRepo
	mchRepo      merchant.IMerchantRepo
	deliveryRepo delivery.IDeliveryRepo
	valRepo      valueobject.IValueRepo
	paymentRepo  payment.IPaymentRepo
	expressRepo  express.IExpressRepo
	mch          merchant.IMerchant
	shipRepo     shipment.IShipmentRepo
	breaker      *wholesaleOrderBreaker
}

func NewOrderManager(cartRepo cart.ICartRepo, mchRepo merchant.IMerchantRepo,
	repo order.IOrderRepo, payRepo payment.IPaymentRepo, productRepo product.IProductRepo,
	goodsRepo item.IGoodsItemRepo, promRepo promotion.IPromotionRepo,
	memberRepo member.IMemberRepo, deliveryRepo delivery.IDeliveryRepo,
	expressRepo express.IExpressRepo, shipRepo shipment.IShipmentRepo,
	valRepo valueobject.IValueRepo) order.IOrderManager {
	return &orderManagerImpl{
		repo:         repo,
		cartRepo:     cartRepo,
		productRepo:  productRepo,
		goodsRepo:    goodsRepo,
		promRepo:     promRepo,
		memberRepo:   memberRepo,
		paymentRepo:  payRepo,
		mchRepo:      mchRepo,
		deliveryRepo: deliveryRepo,
		valRepo:      valRepo,
		expressRepo:  expressRepo,
		shipRepo:     shipRepo,
		breaker:      newWholesaleOrderBreaker(repo),
	}
}

// 在下单前检查购物车
func (t *orderManagerImpl) checkCartForOrder(c cart.ICart) error {
	if c == nil {
		return cart.ErrEmptyShoppingCart
	}
	return c.Check()
}

// 预创建普通订单
func (t *orderManagerImpl) PrepareNormalOrder(c cart.ICart) (order.IOrder, error) {
	err := t.checkCartForOrder(c)
	if err != nil {
		return nil, err
	}
	orderType := order.TRetail
	switch c.Kind() {
	case cart.KRetail:
		orderType = order.TRetail
	case cart.KWholesale:
		orderType = order.TWholesale
	default:
		panic("not support cart kind parse to order")
	}
	val := &order.Order{
		BuyerId:   c.BuyerId(),
		OrderType: int32(orderType),
	}
	o := t.repo.CreateOrder(val)
	if o.Type() != order.TRetail {
		panic("only support normal order")
	}
	io := o.(order.INormalOrder)
	err = io.RequireCart(c)
	io.GetByVendor()
	return o, err
}

// 预创建批发订单
func (o *orderManagerImpl) PrepareWholesaleOrder(c cart.IWholesaleCart) ([]order.IOrder, error) {
	return o.breaker.BreakUp(c)
}

func (t *orderManagerImpl) GetFreeOrderNo(vendorId int32) string {
	return t.repo.GetFreeOrderNo(vendorId)
}

// 生成支付单
func (t *orderManagerImpl) createPaymentOrder(m member.IMember,
	o order.IOrder) payment.IPaymentOrder {
	if o.Type() != order.TRetail {
		panic("not support order type")
	}

	val := o.Complex()
	v := &payment.PaymentOrder{
		BuyUser:     m.GetAggregateRootId(),
		PaymentUser: m.GetAggregateRootId(),
		VendorId:    0,
		OrderId:     0,
		Type:        payment.TypeShopping,
		// 支付单金额
		TotalFee: val.FinalAmount,
		// 余额抵扣
		BalanceDiscount: 0,
		// 积分抵扣
		IntegralDiscount: 0,
		// 系统支付抵扣金额
		SystemDiscount: 0,
		// 优惠券金额
		CouponDiscount: 0,
		// 立减金额
		SubAmount: 0,
		// 调整的金额
		AdjustmentAmount: 0,
		// 支付选项
		PaymentOptFlag: payment.OptPerm,
		// 支付方式
		PaymentSign: enum.PaymentOnlinePay,
		//创建时间
		CreateTime: time.Now().Unix(),
		// 在线支付的交易单号
		OuterNo: "",
		//支付时间
		PaidTime: 0,
		// 状态:  0为未付款，1为已付款，2为已取消
		State: payment.StateAwaitingPayment,
	}
	v.FinalAmount = v.TotalFee - v.SubAmount - v.SystemDiscount -
		v.IntegralDiscount - v.BalanceDiscount
	return t.paymentRepo.CreatePaymentOrder(v)
}

// 应用优惠券
func (t *orderManagerImpl) applyCoupon(m member.IMember, o order.IOrder,
	py payment.IPaymentOrder, couponCode string) error {
	if o.Type() != order.TRetail {
		return errors.New("不支持优惠券")
	}
	io := o.(order.INormalOrder)
	po := py.GetValue()
	cp := t.promRepo.GetCouponByCode(
		m.GetAggregateRootId(), couponCode)
	// 如果优惠券不存在
	if cp == nil {
		return errors.New("优惠券无效")
	}
	// 获取优惠券
	coupon := cp.(promotion.ICouponPromotion)
	result, err := coupon.CanUse(m, po.TotalFee)
	if result {
		if coupon.CanTake() {
			_, err = coupon.GetTake(m.GetAggregateRootId())
			//如果未占用，则占用
			if err != nil {
				err = coupon.Take(m.GetAggregateRootId())
			}
		} else {
			_, err = coupon.GetBind(m.GetAggregateRootId())
		}
		if err != nil {
			domain.HandleError(err, "domain")
			err = errors.New("优惠券无效")
		} else {
			//应用优惠券
			if err = io.ApplyCoupon(coupon); err == nil {
				_, err = py.CouponDiscount(coupon)
			}
		}
	}
	return err
}

func (t *orderManagerImpl) SubmitOrder(c cart.ICart, addressId int32,
	couponCode string, useBalanceDiscount bool) (order.IOrder, error) {
	o, err := t.PrepareNormalOrder(c)
	if err == nil {
		if o.Type() != order.TRetail {
			panic("only support retail cart!")
		}
		io := o.(order.INormalOrder)
		err = io.SetAddress(addressId)
		if err != nil {
			return o, err
		}
		// 更新默认收货地址为本地使用地址
		o.Buyer().Profile().SetDefaultAddress(addressId)

		err = o.Submit()
		buyer := o.Buyer()
		if err == nil {
			if c.Kind() != cart.KRetail {
				panic("购物车非零售")
			}
			rc := c.(cart.IRetailCart)
			cv := rc.GetValue()

			py := io.GetPaymentOrder()
			// 设置支付方式
			cv.PaymentOpt = enum.PaymentOnlinePay
			if err = py.SetPaymentSign(cv.PaymentOpt); err != nil {
				return o, err
			}
			// 使用优惠码
			if len(couponCode) != 0 {
				err = t.applyCoupon(buyer, o, py, couponCode)
				if err != nil {
					return o, err
				}
			}
			// 使用余额抵扣
			if useBalanceDiscount {
				err = py.BalanceDiscount("")
			}
		}
	}
	return o, err
}

// 根据订单编号获取订单
func (t *orderManagerImpl) GetOrderById(orderId int64) order.IOrder {
	val := t.repo.GetOrder("id=?", orderId)
	if val != nil {
		return t.repo.CreateOrder(val)
	}
	return nil
}

// 根据订单号获取订单
func (t *orderManagerImpl) GetOrderByNo(orderNo string) order.IOrder {
	val := t.repo.GetOrder("order_no=?", orderNo)
	if val != nil {
		return t.repo.CreateOrder(val)
	}
	return nil
}

// 接收在线交易支付的通知，不主动调用
func (t *orderManagerImpl) ReceiveNotifyOfOnlineTrade(orderId int64) error {
	o := t.GetOrderById(orderId)
	if o == nil {
		return order.ErrNoSuchOrder
	}
	if o.Type() != order.TRetail {
		panic("unknown order type")
	}
	io := o.(order.INormalOrder)
	return io.OnlinePaymentTradeFinish()
}

// 获取子订单
func (t *orderManagerImpl) GetSubOrder(id int64) order.ISubOrder {
	if v := t.repo.GetSubOrder(id); v != nil {
		return t.repo.CreateNormalSubOrder(v)
	}
	return nil
}
