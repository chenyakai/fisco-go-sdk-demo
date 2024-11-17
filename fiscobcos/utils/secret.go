package utils

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// GeneratePriKey  生成私钥
func GeneratePriKey() (*ecdsa.PrivateKey, error) {
	//SDK发送交易需要一个外部账户，导入go-sdk的`crypto`包，该包提供用于生成随机私钥的`GenerateKey`方法：
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//然后我们可以通过导入golang`crypto/ecdsa`包并使用`FromECDSA`方法将其转换为字节：
	privateKeyBytes := crypto.FromECDSA(privateKey)

	//我们现在可以使用go-sdk的`common/hexutil`包将它转换为十六进制字符串，该包提供了一个带有字节切片的`Encode`方法。 然后我们在十六进制编码之后删除“0x”。
	fmt.Println("Figure PrivateKey: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"
	//这就是`用于签署交易的私钥，将被视为密码，永远不应该被共享给别人`。
	return privateKey, nil
}

// FigurePublicKey 根据私钥计算公钥
func FigurePublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, bool) {
	//由于公钥是从私钥派生的，加密私钥具有一个返回公钥的`Public`方法：
	publicKey := privateKey.Public()

	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了`0x`和前2个字符`04`，它始终是EC前缀，不是必需的。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, ok
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Figure PublicKey: ", hexutil.Encode(publicKeyBytes)[4:]) // publicKey in hex without "0x"
	return publicKeyECDSA, true
}

// FiguredAddress 根据公钥计算地址
func FiguredAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	//现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。 加密包里有一个`PubkeyToAddress`方法，它接受一个ECDSA公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println("address: ", strings.ToLower(address)) // account address
	return address
}
