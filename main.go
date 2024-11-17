package main

import (
	"crypto/ecdsa"
	"fisco-go-sdk-demo/core"
	"fisco-go-sdk-demo/fiscobcos/utils"
	contract "fisco-go-sdk-demo/models/contracts"
	"fmt"
	"math/big"
	"time"
)

const (
	Test_HelloWorld   = "contract1"
	Test_Announcement = "contract2"
)

func TestHelloWorld(key *ecdsa.PrivateKey) {
	//获取string类型
	s1 := new(string)
	utils.SendCallByKey(Test_HelloWorld, "get", key, s1)
	//更改string类型
	transaction := utils.SendTransactionByKey(Test_HelloWorld, "set", key, "Hello,FISCO-BCOS")
	//获取string类型
	s2 := new(string)
	utils.SendCallByKey(Test_HelloWorld, "get", key, s2)
	fmt.Println("SendCall1==>", *s1)
	fmt.Println("transaction2==>", transaction)
	fmt.Println("SendCall12==>", *s2)
}
func main() {
	core.InitConf()
	core.InitClient()
	core.InitSession(Test_HelloWorld)

	//将复制的私钥放到这
	privateKey := "-----BEGIN PRIVATE KEY-----\nMIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQgmqbKrcJgMK6/fVoAuhgV\nbRsg5JCoS+4HD4eFj4qn5f6gBwYFK4EEAAqhRANCAATan9lZU4g6+DXBScOZ0X5U\npkuMg+CAPi85MijzzFof7cY8NvrZX+fy5hrOw5SRnKHnyw4VQYbFzpaqU7L84Pa3\n-----END PRIVATE KEY-----"
	//将这个(pem格式的)私钥转换成 *ecdsa.PrivateKey格式的私钥
	ecdsaKey := utils.PemConvertEcdsa(privateKey)
	//计算公钥
	publicKey, _ := utils.FigurePublicKey(ecdsaKey)
	//计算地址
	address := utils.FiguredAddress(publicKey)
	fmt.Println("FiguredAddress==>", address)
	TestHelloWorld(ecdsaKey)
}

func TestAnnouncement(key *ecdsa.PrivateKey) {
	//addAnouncement  添加公告
	now := time.Now()
	timestamp := now.Unix() // 获取当前时间的Unix时间戳
	fmt.Println("当前Unix时间戳:", timestamp)

	signtime := new(big.Int)
	signtime.SetInt64(timestamp)

	addAnouncementReturn := utils.SendTransactionByKey(Test_Announcement, "addAnouncement", key, "10006", "userPK6", "nounce6", "message6", "cipher6", "attachment6", signtime)
	fmt.Println("addAnouncementReturn=>", addAnouncementReturn)

	//announcements 查看公告
	index := new(big.Int)
	index.SetInt64(6)

	out := new(contract.Output)

	utils.SendCallByKey(Test_Announcement, "announcements", key, out, index)
	fmt.Println("public method announcementsReturn =>", *out)

	//listAnouncement (分页获取数据)
	ans := new([]contract.Output)
	byUserID := true //是否使用UserID，来进行查询
	userID := "10005"

	byTimeRange := true //是否使用时间范围来查询
	startTime := new(big.Int)
	startTime.SetInt64(0) //设置起始时间

	end_now := time.Now()
	end_imestamp := end_now.Unix() // 获取当前时间的Unix时间戳
	fmt.Println("当前Unix时间戳:", end_imestamp)

	endTime := new(big.Int)
	endTime.SetInt64(end_imestamp) //设置结束时间
	pageSize := new(big.Int)       //每页大小
	pageSize.SetInt64(10)
	startIndex := new(big.Int) //起始索引
	startIndex.SetInt64(0)

	utils.SendCallByKey(Test_Announcement, "listAnouncement", key, ans, byUserID, userID, byTimeRange, startTime, endTime, pageSize, startIndex)
	fmt.Println("listAnouncementReturn Length =>", len(*ans))
	fmt.Println("listAnouncementReturn =>", *ans)
}
