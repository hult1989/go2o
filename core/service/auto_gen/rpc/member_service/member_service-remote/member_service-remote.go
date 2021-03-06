// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go2o/core/service/auto_gen/rpc/member_service"
	"go2o/core/service/auto_gen/rpc/message_service"
	"go2o/core/service/auto_gen/rpc/ttype"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var _ = ttype.GoUnusedProtection__
var _ = message_service.GoUnusedProtection__
var _ = member_service.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Result RegisterMemberV2(string user, string pwd, i32 flag, string name, string phone, string email, string avatar,  extend)")
	fmt.Fprintln(os.Stderr, "  Result CheckLogin(string user, string pwd, bool update)")
	fmt.Fprintln(os.Stderr, "  Result CheckTradePwd(i64 memberId, string tradePwd)")
	fmt.Fprintln(os.Stderr, "  i64 SwapMemberId(ECredentials cred, string value)")
	fmt.Fprintln(os.Stderr, "   MemberLevelList()")
	fmt.Fprintln(os.Stderr, "  STrustedInfo GetTrustInfo(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result SubmitTrustInfo(i64 memberId, STrustedInfo info)")
	fmt.Fprintln(os.Stderr, "  Result ReviewTrustedInfo(i64 memberId, bool reviewPass, string remark)")
	fmt.Fprintln(os.Stderr, "  SMemberLevel GetMemberLevel(i32 id)")
	fmt.Fprintln(os.Stderr, "  Result SaveMemberLevel(SMemberLevel level)")
	fmt.Fprintln(os.Stderr, "  SMemberLevel GetLevelBySign(string sign)")
	fmt.Fprintln(os.Stderr, "  SMember GetMember(i64 id)")
	fmt.Fprintln(os.Stderr, "  SMember GetMemberByUser(string user)")
	fmt.Fprintln(os.Stderr, "  SProfile GetProfile(i64 id)")
	fmt.Fprintln(os.Stderr, "  Result Active(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result Lock(i64 memberId, i32 minutes, string remark)")
	fmt.Fprintln(os.Stderr, "  Result Unlock(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result GrantFlag(i64 memberId, i32 flag)")
	fmt.Fprintln(os.Stderr, "  SComplexMember Complex(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result SendCode(i64 memberId, string operation, EMessageChannel msgType)")
	fmt.Fprintln(os.Stderr, "  Result CompareCode(i64 memberId, string code)")
	fmt.Fprintln(os.Stderr, "   ReceiptsCodes(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result SaveReceiptsCode(i64 memberId, SReceiptsCode code)")
	fmt.Fprintln(os.Stderr, "   Bankcards(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result SaveBankcard(i64 memberId, SBankcard card)")
	fmt.Fprintln(os.Stderr, "  Result CheckProfileComplete(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  SMemberLevelInfo MemberLevelInfo(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result UpdateLevel(i64 memberId, i32 level, bool review, i64 paymentOrderId)")
	fmt.Fprintln(os.Stderr, "  Result ChangePhone(i64 memberId, string phone)")
	fmt.Fprintln(os.Stderr, "  Result ChangeUser(i64 memberId, string usr)")
	fmt.Fprintln(os.Stderr, "  Result ModifyPwd(i64 memberId, string old, string pwd)")
	fmt.Fprintln(os.Stderr, "  Result ModifyTradePwd(i64 memberId, string old, string pwd)")
	fmt.Fprintln(os.Stderr, "   OrdersQuantity(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  Result Premium(i64 memberId, i32 v, i64 expires)")
	fmt.Fprintln(os.Stderr, "  string GetToken(i64 memberId, bool reset)")
	fmt.Fprintln(os.Stderr, "  bool CheckToken(i64 memberId, string token)")
	fmt.Fprintln(os.Stderr, "  void RemoveToken(i64 memberId)")
	fmt.Fprintln(os.Stderr, "   GetAddressList(i64 memberId)")
	fmt.Fprintln(os.Stderr, "  SAddress GetAddress(i64 memberId, i64 addrId)")
	fmt.Fprintln(os.Stderr, "  SAccount GetAccount(i64 memberId)")
	fmt.Fprintln(os.Stderr, "   InviterArray(i64 memberId, i32 depth)")
	fmt.Fprintln(os.Stderr, "  i32 InviteMembersQuantity(i64 memberId, i32 depth)")
	fmt.Fprintln(os.Stderr, "  i32 QueryInviteQuantity(i64 memberId,  data)")
	fmt.Fprintln(os.Stderr, "   QueryInviteArray(i64 memberId,  data)")
	fmt.Fprintln(os.Stderr, "  Result AccountCharge(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
	fmt.Fprintln(os.Stderr, "  Result AccountConsume(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
	fmt.Fprintln(os.Stderr, "  Result AccountDiscount(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
	fmt.Fprintln(os.Stderr, "  Result AccountRefund(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
	fmt.Fprintln(os.Stderr, "  Result AccountAdjust(i64 memberId, i32 account, i32 value, i64 relateUser, string remark)")
	fmt.Fprintln(os.Stderr, "  Result B4EAuth(i64 memberId, string action,  data)")
	fmt.Fprintln(os.Stderr, "  SPagingResult PagingAccountLog(i64 memberId, i32 accountType, SPagingParams params)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
	parts := strings.Split(value, ": ")
	if len(parts) != 2 {
		return fmt.Errorf("header should be of format 'Key: Value'")
	}
	h[parts[0]] = parts[1]
	return nil
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	headers := make(httpHeaders)
	var parsedUrl *url.URL
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
	flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
	flag.Parse()

	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
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
		if len(headers) > 0 {
			httptrans := trans.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
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
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := member_service.NewMemberServiceClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "RegisterMemberV2":
		if flag.NArg()-1 != 8 {
			fmt.Fprintln(os.Stderr, "RegisterMemberV2 requires 8 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		tmp2, err122 := (strconv.Atoi(flag.Arg(3)))
		if err122 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5 := flag.Arg(6)
		value5 := argvalue5
		argvalue6 := flag.Arg(7)
		value6 := argvalue6
		arg127 := flag.Arg(8)
		mbTrans128 := thrift.NewTMemoryBufferLen(len(arg127))
		defer mbTrans128.Close()
		_, err129 := mbTrans128.WriteString(arg127)
		if err129 != nil {
			Usage()
			return
		}
		factory130 := thrift.NewTJSONProtocolFactory()
		jsProt131 := factory130.GetProtocol(mbTrans128)
		containerStruct7 := member_service.NewMemberServiceRegisterMemberV2Args()
		err132 := containerStruct7.ReadField8(jsProt131)
		if err132 != nil {
			Usage()
			return
		}
		argvalue7 := containerStruct7.Extend
		value7 := argvalue7
		fmt.Print(client.RegisterMemberV2(context.Background(), value0, value1, value2, value3, value4, value5, value6, value7))
		fmt.Print("\n")
		break
	case "CheckLogin":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "CheckLogin requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.CheckLogin(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "CheckTradePwd":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CheckTradePwd requires 2 args")
			flag.Usage()
		}
		argvalue0, err136 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err136 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.CheckTradePwd(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "SwapMemberId":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SwapMemberId requires 2 args")
			flag.Usage()
		}
		tmp0, err := (strconv.Atoi(flag.Arg(1)))
		if err != nil {
			Usage()
			return
		}
		argvalue0 := member_service.ECredentials(tmp0)
		value0 := member_service.ECredentials(argvalue0)
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SwapMemberId(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "MemberLevelList":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "MemberLevelList requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.MemberLevelList(context.Background()))
		fmt.Print("\n")
		break
	case "GetTrustInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTrustInfo requires 1 args")
			flag.Usage()
		}
		argvalue0, err139 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err139 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTrustInfo(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SubmitTrustInfo":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SubmitTrustInfo requires 2 args")
			flag.Usage()
		}
		argvalue0, err140 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err140 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg141 := flag.Arg(2)
		mbTrans142 := thrift.NewTMemoryBufferLen(len(arg141))
		defer mbTrans142.Close()
		_, err143 := mbTrans142.WriteString(arg141)
		if err143 != nil {
			Usage()
			return
		}
		factory144 := thrift.NewTJSONProtocolFactory()
		jsProt145 := factory144.GetProtocol(mbTrans142)
		argvalue1 := member_service.NewSTrustedInfo()
		err146 := argvalue1.Read(jsProt145)
		if err146 != nil {
			Usage()
			return
		}
		value1 := member_service.STrustedInfo(argvalue1)
		fmt.Print(client.SubmitTrustInfo(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "ReviewTrustedInfo":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "ReviewTrustedInfo requires 3 args")
			flag.Usage()
		}
		argvalue0, err147 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err147 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.ReviewTrustedInfo(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "GetMemberLevel":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMemberLevel requires 1 args")
			flag.Usage()
		}
		tmp0, err150 := (strconv.Atoi(flag.Arg(1)))
		if err150 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetMemberLevel(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SaveMemberLevel":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SaveMemberLevel requires 1 args")
			flag.Usage()
		}
		arg151 := flag.Arg(1)
		mbTrans152 := thrift.NewTMemoryBufferLen(len(arg151))
		defer mbTrans152.Close()
		_, err153 := mbTrans152.WriteString(arg151)
		if err153 != nil {
			Usage()
			return
		}
		factory154 := thrift.NewTJSONProtocolFactory()
		jsProt155 := factory154.GetProtocol(mbTrans152)
		argvalue0 := member_service.NewSMemberLevel()
		err156 := argvalue0.Read(jsProt155)
		if err156 != nil {
			Usage()
			return
		}
		value0 := member_service.SMemberLevel(argvalue0)
		fmt.Print(client.SaveMemberLevel(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetLevelBySign":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetLevelBySign requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetLevelBySign(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetMember":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMember requires 1 args")
			flag.Usage()
		}
		argvalue0, err158 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err158 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetMember(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetMemberByUser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMemberByUser requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetMemberByUser(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetProfile":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetProfile requires 1 args")
			flag.Usage()
		}
		argvalue0, err160 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err160 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetProfile(context.Background(), value0))
		fmt.Print("\n")
		break
	case "Active":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Active requires 1 args")
			flag.Usage()
		}
		argvalue0, err161 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err161 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Active(context.Background(), value0))
		fmt.Print("\n")
		break
	case "Lock":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "Lock requires 3 args")
			flag.Usage()
		}
		argvalue0, err162 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err162 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err163 := (strconv.Atoi(flag.Arg(2)))
		if err163 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.Lock(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "Unlock":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Unlock requires 1 args")
			flag.Usage()
		}
		argvalue0, err165 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err165 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Unlock(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GrantFlag":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GrantFlag requires 2 args")
			flag.Usage()
		}
		argvalue0, err166 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err166 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err167 := (strconv.Atoi(flag.Arg(2)))
		if err167 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.GrantFlag(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "Complex":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Complex requires 1 args")
			flag.Usage()
		}
		argvalue0, err168 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err168 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Complex(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SendCode":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "SendCode requires 3 args")
			flag.Usage()
		}
		argvalue0, err169 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err169 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		tmp2, err := (strconv.Atoi(flag.Arg(3)))
		if err != nil {
			Usage()
			return
		}
		argvalue2 := member_service.EMessageChannel(tmp2)
		value2 := argvalue2
		fmt.Print(client.SendCode(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "CompareCode":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CompareCode requires 2 args")
			flag.Usage()
		}
		argvalue0, err171 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err171 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.CompareCode(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "ReceiptsCodes":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ReceiptsCodes requires 1 args")
			flag.Usage()
		}
		argvalue0, err173 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err173 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ReceiptsCodes(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SaveReceiptsCode":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SaveReceiptsCode requires 2 args")
			flag.Usage()
		}
		argvalue0, err174 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err174 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg175 := flag.Arg(2)
		mbTrans176 := thrift.NewTMemoryBufferLen(len(arg175))
		defer mbTrans176.Close()
		_, err177 := mbTrans176.WriteString(arg175)
		if err177 != nil {
			Usage()
			return
		}
		factory178 := thrift.NewTJSONProtocolFactory()
		jsProt179 := factory178.GetProtocol(mbTrans176)
		argvalue1 := member_service.NewSReceiptsCode()
		err180 := argvalue1.Read(jsProt179)
		if err180 != nil {
			Usage()
			return
		}
		value1 := member_service.SReceiptsCode(argvalue1)
		fmt.Print(client.SaveReceiptsCode(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "Bankcards":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Bankcards requires 1 args")
			flag.Usage()
		}
		argvalue0, err181 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err181 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Bankcards(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SaveBankcard":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SaveBankcard requires 2 args")
			flag.Usage()
		}
		argvalue0, err182 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err182 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg183 := flag.Arg(2)
		mbTrans184 := thrift.NewTMemoryBufferLen(len(arg183))
		defer mbTrans184.Close()
		_, err185 := mbTrans184.WriteString(arg183)
		if err185 != nil {
			Usage()
			return
		}
		factory186 := thrift.NewTJSONProtocolFactory()
		jsProt187 := factory186.GetProtocol(mbTrans184)
		argvalue1 := member_service.NewSBankcard()
		err188 := argvalue1.Read(jsProt187)
		if err188 != nil {
			Usage()
			return
		}
		value1 := member_service.SBankcard(argvalue1)
		fmt.Print(client.SaveBankcard(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "CheckProfileComplete":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CheckProfileComplete requires 1 args")
			flag.Usage()
		}
		argvalue0, err189 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err189 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CheckProfileComplete(context.Background(), value0))
		fmt.Print("\n")
		break
	case "MemberLevelInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MemberLevelInfo requires 1 args")
			flag.Usage()
		}
		argvalue0, err190 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err190 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MemberLevelInfo(context.Background(), value0))
		fmt.Print("\n")
		break
	case "UpdateLevel":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "UpdateLevel requires 4 args")
			flag.Usage()
		}
		argvalue0, err191 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err191 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err192 := (strconv.Atoi(flag.Arg(2)))
		if err192 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		argvalue3, err194 := (strconv.ParseInt(flag.Arg(4), 10, 64))
		if err194 != nil {
			Usage()
			return
		}
		value3 := argvalue3
		fmt.Print(client.UpdateLevel(context.Background(), value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "ChangePhone":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ChangePhone requires 2 args")
			flag.Usage()
		}
		argvalue0, err195 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err195 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.ChangePhone(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "ChangeUser":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ChangeUser requires 2 args")
			flag.Usage()
		}
		argvalue0, err197 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err197 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.ChangeUser(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "ModifyPwd":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "ModifyPwd requires 3 args")
			flag.Usage()
		}
		argvalue0, err199 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err199 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.ModifyPwd(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "ModifyTradePwd":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "ModifyTradePwd requires 3 args")
			flag.Usage()
		}
		argvalue0, err202 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err202 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.ModifyTradePwd(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "OrdersQuantity":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "OrdersQuantity requires 1 args")
			flag.Usage()
		}
		argvalue0, err205 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err205 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.OrdersQuantity(context.Background(), value0))
		fmt.Print("\n")
		break
	case "Premium":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "Premium requires 3 args")
			flag.Usage()
		}
		argvalue0, err206 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err206 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err207 := (strconv.Atoi(flag.Arg(2)))
		if err207 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2, err208 := (strconv.ParseInt(flag.Arg(3), 10, 64))
		if err208 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.Premium(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "GetToken":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetToken requires 2 args")
			flag.Usage()
		}
		argvalue0, err209 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err209 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetToken(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "CheckToken":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CheckToken requires 2 args")
			flag.Usage()
		}
		argvalue0, err211 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err211 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.CheckToken(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "RemoveToken":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveToken requires 1 args")
			flag.Usage()
		}
		argvalue0, err213 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err213 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RemoveToken(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetAddressList":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAddressList requires 1 args")
			flag.Usage()
		}
		argvalue0, err214 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err214 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetAddressList(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetAddress":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetAddress requires 2 args")
			flag.Usage()
		}
		argvalue0, err215 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err215 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err216 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err216 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.GetAddress(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "GetAccount":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAccount requires 1 args")
			flag.Usage()
		}
		argvalue0, err217 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err217 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetAccount(context.Background(), value0))
		fmt.Print("\n")
		break
	case "InviterArray":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "InviterArray requires 2 args")
			flag.Usage()
		}
		argvalue0, err218 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err218 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err219 := (strconv.Atoi(flag.Arg(2)))
		if err219 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.InviterArray(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "InviteMembersQuantity":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "InviteMembersQuantity requires 2 args")
			flag.Usage()
		}
		argvalue0, err220 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err220 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err221 := (strconv.Atoi(flag.Arg(2)))
		if err221 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.InviteMembersQuantity(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "QueryInviteQuantity":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "QueryInviteQuantity requires 2 args")
			flag.Usage()
		}
		argvalue0, err222 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err222 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg223 := flag.Arg(2)
		mbTrans224 := thrift.NewTMemoryBufferLen(len(arg223))
		defer mbTrans224.Close()
		_, err225 := mbTrans224.WriteString(arg223)
		if err225 != nil {
			Usage()
			return
		}
		factory226 := thrift.NewTJSONProtocolFactory()
		jsProt227 := factory226.GetProtocol(mbTrans224)
		containerStruct1 := member_service.NewMemberServiceQueryInviteQuantityArgs()
		err228 := containerStruct1.ReadField2(jsProt227)
		if err228 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Data
		value1 := argvalue1
		fmt.Print(client.QueryInviteQuantity(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "QueryInviteArray":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "QueryInviteArray requires 2 args")
			flag.Usage()
		}
		argvalue0, err229 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err229 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg230 := flag.Arg(2)
		mbTrans231 := thrift.NewTMemoryBufferLen(len(arg230))
		defer mbTrans231.Close()
		_, err232 := mbTrans231.WriteString(arg230)
		if err232 != nil {
			Usage()
			return
		}
		factory233 := thrift.NewTJSONProtocolFactory()
		jsProt234 := factory233.GetProtocol(mbTrans231)
		containerStruct1 := member_service.NewMemberServiceQueryInviteArrayArgs()
		err235 := containerStruct1.ReadField2(jsProt234)
		if err235 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Data
		value1 := argvalue1
		fmt.Print(client.QueryInviteArray(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "AccountCharge":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "AccountCharge requires 6 args")
			flag.Usage()
		}
		argvalue0, err236 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err236 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err237 := (strconv.Atoi(flag.Arg(2)))
		if err237 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		tmp3, err239 := (strconv.Atoi(flag.Arg(4)))
		if err239 != nil {
			Usage()
			return
		}
		argvalue3 := int32(tmp3)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5 := flag.Arg(6)
		value5 := argvalue5
		fmt.Print(client.AccountCharge(context.Background(), value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "AccountConsume":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "AccountConsume requires 6 args")
			flag.Usage()
		}
		argvalue0, err242 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err242 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err243 := (strconv.Atoi(flag.Arg(2)))
		if err243 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		tmp3, err245 := (strconv.Atoi(flag.Arg(4)))
		if err245 != nil {
			Usage()
			return
		}
		argvalue3 := int32(tmp3)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5 := flag.Arg(6)
		value5 := argvalue5
		fmt.Print(client.AccountConsume(context.Background(), value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "AccountDiscount":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "AccountDiscount requires 6 args")
			flag.Usage()
		}
		argvalue0, err248 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err248 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err249 := (strconv.Atoi(flag.Arg(2)))
		if err249 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		tmp3, err251 := (strconv.Atoi(flag.Arg(4)))
		if err251 != nil {
			Usage()
			return
		}
		argvalue3 := int32(tmp3)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5 := flag.Arg(6)
		value5 := argvalue5
		fmt.Print(client.AccountDiscount(context.Background(), value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "AccountRefund":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "AccountRefund requires 6 args")
			flag.Usage()
		}
		argvalue0, err254 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err254 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err255 := (strconv.Atoi(flag.Arg(2)))
		if err255 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		tmp3, err257 := (strconv.Atoi(flag.Arg(4)))
		if err257 != nil {
			Usage()
			return
		}
		argvalue3 := int32(tmp3)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5 := flag.Arg(6)
		value5 := argvalue5
		fmt.Print(client.AccountRefund(context.Background(), value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "AccountAdjust":
		if flag.NArg()-1 != 5 {
			fmt.Fprintln(os.Stderr, "AccountAdjust requires 5 args")
			flag.Usage()
		}
		argvalue0, err260 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err260 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err261 := (strconv.Atoi(flag.Arg(2)))
		if err261 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		tmp2, err262 := (strconv.Atoi(flag.Arg(3)))
		if err262 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		argvalue3, err263 := (strconv.ParseInt(flag.Arg(4), 10, 64))
		if err263 != nil {
			Usage()
			return
		}
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		fmt.Print(client.AccountAdjust(context.Background(), value0, value1, value2, value3, value4))
		fmt.Print("\n")
		break
	case "B4EAuth":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "B4EAuth requires 3 args")
			flag.Usage()
		}
		argvalue0, err265 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err265 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg267 := flag.Arg(3)
		mbTrans268 := thrift.NewTMemoryBufferLen(len(arg267))
		defer mbTrans268.Close()
		_, err269 := mbTrans268.WriteString(arg267)
		if err269 != nil {
			Usage()
			return
		}
		factory270 := thrift.NewTJSONProtocolFactory()
		jsProt271 := factory270.GetProtocol(mbTrans268)
		containerStruct2 := member_service.NewMemberServiceB4EAuthArgs()
		err272 := containerStruct2.ReadField3(jsProt271)
		if err272 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Data
		value2 := argvalue2
		fmt.Print(client.B4EAuth(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "PagingAccountLog":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "PagingAccountLog requires 3 args")
			flag.Usage()
		}
		argvalue0, err273 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err273 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err274 := (strconv.Atoi(flag.Arg(2)))
		if err274 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		arg275 := flag.Arg(3)
		mbTrans276 := thrift.NewTMemoryBufferLen(len(arg275))
		defer mbTrans276.Close()
		_, err277 := mbTrans276.WriteString(arg275)
		if err277 != nil {
			Usage()
			return
		}
		factory278 := thrift.NewTJSONProtocolFactory()
		jsProt279 := factory278.GetProtocol(mbTrans276)
		argvalue2 := ttype.NewSPagingParams()
		err280 := argvalue2.Read(jsProt279)
		if err280 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.PagingAccountLog(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
