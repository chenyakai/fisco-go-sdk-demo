package main

//
//import (
//	"crypto/ecdsa"
//	"fisco-go-sdk-demo/core"
//	"fisco-go-sdk-demo/fiscobcos/utils"
//	"fisco-go-sdk-demo/global"
//	contract "fisco-go-sdk-demo/models/contracts"
//	"fmt"
//	"math/big"
//	"time"
//)
//
//func HelloWorld(sessionAbi *string, key *ecdsa.PrivateKey) {
//	//获取string类型
//	s1 := new(string)
//	//utils.SendCall("contract1", "get", s1)
//
//	utils.SendCallByKey("contract1", "get", key, s1)
//
//	//更改string类型
//	//transaction := utils.SendTransaction("contract1", "set", *sessionAbi, "Hello,FISCO-BCOS")
//	transaction := utils.SendTransactionByKey("contract1", "set", key, *sessionAbi, "Hello,FISCO-BCOS")
//
//	//获取string类型
//	s2 := new(string)
//	//utils.SendCall("contract1", "get", s2)
//	utils.SendCallByKey("contract1", "get", key, s2)
//
//	fmt.Println("SendCall1==>", *s1)
//	fmt.Println("transaction2==>", transaction)
//	fmt.Println("SendCall12==>", *s2)
//}
//
//func main() {
//
//	core.InitConf()
//	core.InitClient()
//	//key, err := utils.GeneratePriKey()
//	//if err != nil {
//	//	fmt.Println("GeneratePriKey===>", err)
//	//}
//
//	key := "10938885802782121922763370745052833944275575189324663687624187374667059855501"
//	ecdsaKey := utils.DeConvertEcdsa(key)
//
//	sessionAbi1 := core.InitSession("contract1", global.Config.Contract["contract1"].Address, global.Config.Contract["contract1"].Abi)
//	HelloWorld(sessionAbi1, ecdsaKey)
//	//address := global.GoSdk.Client.GetTransactOpts().From.Hex()
//	//address := global.GoSdk.Client.GetTransactOpts()
//	//fmt.Println("address==>", address)
//
//	publicKey, b := utils.FigurePublicKey(ecdsaKey)
//	if !b {
//		fmt.Println("FigurePublicKey Failed")
//	}
//	figuredAddress := utils.FiguredAddress(publicKey)
//	fmt.Println("figuredAddress==>", figuredAddress)
//	//if strings.EqualFold(figuredAddress, address) {
//	//	fmt.Println("figuredAddress，address地址一致")
//	//} else {
//	//	fmt.Println("figuredAddress，address地址不一致")
//	//}
//
//	SessionAbi2 := core.InitSession("contract2", global.Config.Contract["contract2"].Address, global.Config.Contract["contract2"].Abi)
//	Announcement(SessionAbi2, ecdsaKey)
//
//	//获取string类型
//	//s1 := new(string)
//	//utils.SendTransactionByKey("contract1", "set", key, *sessionAbi1, "HelloWorld")
//	//fmt.Println("SendCall12==>", *s1)
//
//	//announcements
//	index := new(big.Int)
//	index.SetInt64(6)
//
//	out := new(contract.Output)
//
//	utils.SendCallByKey("contract2", "announcements", ecdsaKey, out, index)
//	fmt.Println("public method announcementsReturn =>", *out)
//}
//
//func Announcement(sessionAbi *string, key *ecdsa.PrivateKey) {
//	//addAnouncement
//
//	now := time.Now()
//	timestamp := now.Unix() // 获取当前时间的Unix时间戳
//	fmt.Println("当前Unix时间戳:", timestamp)
//
//	signtime := new(big.Int)
//	signtime.SetInt64(timestamp)
//
//	//addAnouncementReturn := utils.SendTransaction("contract2", "addAnouncement", *sessionAbi, "10006", "userPK6", "nounce6", "message6", "cipher6", "attachment6", signtime)
//	addAnouncementReturn := utils.SendTransactionByKey("contract2", "addAnouncement", key, *sessionAbi, "10006", "userPK6", "nounce6", "message6", "cipher6", "attachment6", signtime)
//	fmt.Println("addAnouncementReturn=>", addAnouncementReturn)
//
//	//announcements
//
//	index := new(big.Int)
//	index.SetInt64(6)
//
//	out := new(contract.Output)
//
//	//utils.SendCall("contract2", "announcements", out, index)
//	utils.SendCallByKey("contract2", "announcements", key, out, index)
//	fmt.Println("public method announcementsReturn =>", *out)
//
//	//listAnouncement (分页获取数据)
//	ans := new([]contract.Output)
//	byUserID := false //是否使用UserID，来进行查询
//	userID := "10005"
//
//	byTimeRange := true //是否使用时间范围来查询
//	startTime := new(big.Int)
//	startTime.SetInt64(0) //设置起始时间
//
//	end_now := time.Now()
//	end_imestamp := end_now.Unix() // 获取当前时间的Unix时间戳
//	fmt.Println("当前Unix时间戳:", end_imestamp)
//
//	endTime := new(big.Int)
//	endTime.SetInt64(end_imestamp) //设置结束时间
//	pageSize := new(big.Int)       //每页大小
//	pageSize.SetInt64(10)
//	startIndex := new(big.Int) //起始索引
//	startIndex.SetInt64(0)
//
//	//utils.SendCall("contract2", "listAnouncement", ans, byUserID, userID, byTimeRange, startTime, endTime, pageSize, startIndex)
//	utils.SendCallByKey("contract2", "listAnouncement", key, ans, byUserID, userID, byTimeRange, startTime, endTime, pageSize, startIndex)
//	fmt.Println("listAnouncementReturn Length =>", len(*ans))
//	fmt.Println("listAnouncementReturn =>", *ans)
//}
