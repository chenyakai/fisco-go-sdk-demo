package main

func main() {
	//十六进制私钥字符串。（测试：获取*ecdsa.PrivateKey类型私钥，计算公钥，地址与实际值，是否相同。）
	//s := "9663ee21a52c11159ba9cada57ffc1ff700a1ca383cef133cb6cc7772e8f73e4"
	//ecdsaKey := utils.HexConvertEcdsa(s)
	//key, b := utils.FigurePublicKey(ecdsaKey)
	//if !b {
	//	fmt.Println("FigurePublicKey invoke Failed==>", b)
	//}
	////fmt.Println("ecdsa_key.PublicKey==>", *key)
	//fmt.Println("ecdsa_key.Address==>", utils.FiguredAddress(key))

	//pem私钥字符串。（测试：获取*ecdsa.PrivateKey类型私钥，计算公钥，地址与实际值，是否相同。）
	//s := "-----BEGIN PRIVATE KEY-----\nMIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQglmPuIaUsERWbqcraV//B\n/3AKHKODzvEzy2zHdy6Pc+SgBwYFK4EEAAqhRANCAAQNJwd4sQ9LHzMnQNPsnf+e\n25++YJlEcrw71xigJYNQWHSCwXpfRft8w+EItTzhq3Y3Wcatvv8cU08TleLzgt/y\n-----END PRIVATE KEY-----"
	//
	//ecdsaKey := utils.PemConvertEcdsa(s)
	//key, b := utils.FigurePublicKey(ecdsaKey)
	//if !b {
	//	fmt.Println("FigurePublicKey invoke Failed==>", b)
	//}
	////fmt.Println("ecdsa_key.PublicKey==>", *key)
	//fmt.Println("ecdsa_key.Address==>", strings.ToLower(utils.FiguredAddress(key)))

	//十进制私钥字符串
	//s := "68023488669606450603805102916363711927531764781668974227744976329694193349604"
	//
	//ecdsaKey := utils.DeConvertEcdsa(s)
	//key, b := utils.FigurePublicKey(ecdsaKey)
	//if !b {
	//	fmt.Println("FigurePublicKey invoke Failed==>", b)
	//}
	////fmt.Println("ecdsa_key.PublicKey==>", *key)
	//fmt.Println("ecdsa_key.Address==>", strings.ToLower(utils.FiguredAddress(key)))

}
