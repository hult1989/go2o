// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"define"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Result WholesaleCartV1(i64 memberId, string action,  data)")
	fmt.Fprintln(os.Stderr, "  Result RetailCartV1(i64 memberId, string action,  data)")
	fmt.Fprintln(os.Stderr, "   SubmitOrderV1(i64 buyerId, i32 cartType,  data)")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetOrder(string order_no, bool sub_order)")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetOrderAndItems(string order_no, bool sub_order)")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetSubOrder(i64 id)")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetSubOrderByNo(string orderNo)")
	fmt.Fprintln(os.Stderr, "   GetSubOrderItems(i64 subOrderId)")
	fmt.Fprintln(os.Stderr, "  Result64 SubmitTradeOrder(ComplexOrder o, double rate)")
	fmt.Fprintln(os.Stderr, "  Result64 TradeOrderCashPay(i64 orderId)")
	fmt.Fprintln(os.Stderr, "  Result64 TradeOrderUpdateTicket(i64 orderId, string img)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := define.NewSaleServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "WholesaleCartV1":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "WholesaleCartV1 requires 3 args")
			flag.Usage()
		}
		argvalue0, err251 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err251 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg253 := flag.Arg(3)
		mbTrans254 := thrift.NewTMemoryBufferLen(len(arg253))
		defer mbTrans254.Close()
		_, err255 := mbTrans254.WriteString(arg253)
		if err255 != nil {
			Usage()
			return
		}
		factory256 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt257 := factory256.GetProtocol(mbTrans254)
		containerStruct2 := define.NewSaleServiceWholesaleCartV1Args()
		err258 := containerStruct2.ReadField3(jsProt257)
		if err258 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Data
		value2 := argvalue2
		fmt.Print(client.WholesaleCartV1(value0, value1, value2))
		fmt.Print("\n")
		break
	case "RetailCartV1":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RetailCartV1 requires 3 args")
			flag.Usage()
		}
		argvalue0, err259 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err259 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg261 := flag.Arg(3)
		mbTrans262 := thrift.NewTMemoryBufferLen(len(arg261))
		defer mbTrans262.Close()
		_, err263 := mbTrans262.WriteString(arg261)
		if err263 != nil {
			Usage()
			return
		}
		factory264 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt265 := factory264.GetProtocol(mbTrans262)
		containerStruct2 := define.NewSaleServiceRetailCartV1Args()
		err266 := containerStruct2.ReadField3(jsProt265)
		if err266 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Data
		value2 := argvalue2
		fmt.Print(client.RetailCartV1(value0, value1, value2))
		fmt.Print("\n")
		break
	case "SubmitOrderV1":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "SubmitOrderV1 requires 3 args")
			flag.Usage()
		}
		argvalue0, err267 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err267 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err268 := (strconv.Atoi(flag.Arg(2)))
		if err268 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		arg269 := flag.Arg(3)
		mbTrans270 := thrift.NewTMemoryBufferLen(len(arg269))
		defer mbTrans270.Close()
		_, err271 := mbTrans270.WriteString(arg269)
		if err271 != nil {
			Usage()
			return
		}
		factory272 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt273 := factory272.GetProtocol(mbTrans270)
		containerStruct2 := define.NewSaleServiceSubmitOrderV1Args()
		err274 := containerStruct2.ReadField3(jsProt273)
		if err274 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Data
		value2 := argvalue2
		fmt.Print(client.SubmitOrderV1(value0, value1, value2))
		fmt.Print("\n")
		break
	case "GetOrder":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetOrder requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetOrder(value0, value1))
		fmt.Print("\n")
		break
	case "GetOrderAndItems":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetOrderAndItems requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetOrderAndItems(value0, value1))
		fmt.Print("\n")
		break
	case "GetSubOrder":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSubOrder requires 1 args")
			flag.Usage()
		}
		argvalue0, err279 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err279 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetSubOrder(value0))
		fmt.Print("\n")
		break
	case "GetSubOrderByNo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSubOrderByNo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetSubOrderByNo(value0))
		fmt.Print("\n")
		break
	case "GetSubOrderItems":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSubOrderItems requires 1 args")
			flag.Usage()
		}
		argvalue0, err281 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err281 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetSubOrderItems(value0))
		fmt.Print("\n")
		break
	case "SubmitTradeOrder":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SubmitTradeOrder requires 2 args")
			flag.Usage()
		}
		arg282 := flag.Arg(1)
		mbTrans283 := thrift.NewTMemoryBufferLen(len(arg282))
		defer mbTrans283.Close()
		_, err284 := mbTrans283.WriteString(arg282)
		if err284 != nil {
			Usage()
			return
		}
		factory285 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt286 := factory285.GetProtocol(mbTrans283)
		argvalue0 := define.NewComplexOrder()
		err287 := argvalue0.Read(jsProt286)
		if err287 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err288 := (strconv.ParseFloat(flag.Arg(2), 64))
		if err288 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.SubmitTradeOrder(value0, value1))
		fmt.Print("\n")
		break
	case "TradeOrderCashPay":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TradeOrderCashPay requires 1 args")
			flag.Usage()
		}
		argvalue0, err289 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err289 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TradeOrderCashPay(value0))
		fmt.Print("\n")
		break
	case "TradeOrderUpdateTicket":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TradeOrderUpdateTicket requires 2 args")
			flag.Usage()
		}
		argvalue0, err290 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err290 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.TradeOrderUpdateTicket(value0, value1))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
