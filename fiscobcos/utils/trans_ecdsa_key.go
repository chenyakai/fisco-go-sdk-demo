package utils

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"math/big"
)

// 将十六进制私钥转换成ecdsa私钥
func HexConvertEcdsa(key string) *ecdsa.PrivateKey {
	// 将十六进制私钥字符串解码为字节切片
	byteArray, err := hex.DecodeString(key)
	if err != nil {
		fmt.Printf("Error decoding hex string: %v\n", err)
		return nil
	}
	privateKey, ok := crypto.ToECDSA(byteArray)
	if ok != nil {
		fmt.Println("covert invoke crypto.ToECDSA Failed", ok)
	}
	return privateKey
}

// 将十进制私钥转换成ecdsa私钥
func DeConvertEcdsa(key string) *ecdsa.PrivateKey {
	// 将十进制私钥字符串解码为字节切片
	// 使用 math/big 的 NewInt 方法将十进制字符串转换为大整数
	privateKeyInt := new(big.Int)
	var success bool
	privateKeyInt, success = privateKeyInt.SetString(key, 10) // 10表示十进制
	if !success {
		fmt.Println("Error parsing private key")
		return nil
	}

	// 将大整数转换为字节数组
	privateKeyBytes := privateKeyInt.Bytes()
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		fmt.Println("covert invoke crypto.ToECDSA Failed", err)
	}
	return privateKey
}

// 将pem私钥转换成ecdsa私钥

func PemConvertEcdsa(key string) *ecdsa.PrivateKey {
	IsSMCrypto := false //非国密
	keyBytes, curve, err := LoadECPrivateKeyFromPEM(key)
	if err != nil {
		fmt.Errorf("parse private key failed, err: %v", err)
	}
	if IsSMCrypto && curve != "sm2p256v1" {
		fmt.Errorf("smcrypto must use sm2p256v1 private key, but found %s", curve)
	}
	if !IsSMCrypto && curve != "secp256k1" {
		fmt.Errorf("must use secp256k1 private key, but found %s", curve)
	}
	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		logrus.Fatal(err)
	}
	return privateKey
}

// LoadECPrivateKeyFromPEM reads file, divides into key and certificates
func LoadECPrivateKeyFromPEM(key string) ([]byte, string, error) {
	// 移除PEM头尾，只保留Base64编码的部分
	block, _ := pem.Decode([]byte(key))

	if block == nil {
		return nil, "", fmt.Errorf("Failure reading pem from %s", key)
	}
	if block.Type != "PRIVATE KEY" {
		return nil, "", fmt.Errorf("Failure reading private key from %s", key)
	}
	ecPirvateKey, curveName, ok := parsePKCS8ECPrivateKey(block.Bytes)
	if ok != nil {
		return nil, "", fmt.Errorf("Failure reading private key from \"%s\": %s", key, ok)
	}
	return ecPirvateKey, curveName, nil
}

// parseECPrivateKey is a copy of x509.ParseECPrivateKey, supported secp256k1 and sm2p256v1
func parsePKCS8ECPrivateKey(der []byte) (keyHex []byte, curveName string, err error) {
	oidNamedCurveSm2p256v1 := asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 301}
	oidNamedCurveSecp256k1 := asn1.ObjectIdentifier{1, 3, 132, 0, 10}

	oidPublicKeyECDSA := asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
	// AlgorithmIdentifier represents the ASN.1 structure of the same name. See RFC
	// 5280, section 4.1.1.2.
	type AlgorithmIdentifier struct {
		Algorithm  asn1.ObjectIdentifier
		Parameters asn1.RawValue `asn1:"optional"`
	}
	var pkcs8 struct {
		Version    int
		Algo       AlgorithmIdentifier
		PrivateKey []byte
		// optional attributes omitted.
	}
	var privKey struct {
		Version       int
		PrivateKey    []byte
		NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
		PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
	}
	if _, err := asn1.Unmarshal(der, &pkcs8); err != nil {
		return nil, "", errors.New("x509: failed to parse EC private key embedded in PKCS#8: " + err.Error())
	}
	if !pkcs8.Algo.Algorithm.Equal(oidPublicKeyECDSA) {
		return nil, "", fmt.Errorf("x509: PKCS#8 wrapping contained private key with unknown algorithm: %v", pkcs8.Algo.Algorithm)
	}
	bytes := pkcs8.Algo.Parameters.FullBytes
	namedCurveOID := new(asn1.ObjectIdentifier)
	if _, err := asn1.Unmarshal(bytes, namedCurveOID); err != nil {
		namedCurveOID = nil
		return nil, "", fmt.Errorf("parse namedCurveOID failed")
	}
	if _, err := asn1.Unmarshal(pkcs8.PrivateKey, &privKey); err != nil {
		return nil, "", errors.New("x509: failed to parse EC private key: " + err.Error())
	}
	var curveOrder *big.Int

	switch {
	case namedCurveOID.Equal(oidNamedCurveSecp256k1):
		curveName = "secp256k1"
		curveOrder, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	case namedCurveOID.Equal(oidNamedCurveSm2p256v1):
		curveName = "sm2p256v1"
		curveOrder, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	default:
		return nil, "", fmt.Errorf("unknown namedCurveOID:%+v", namedCurveOID)
	}

	k := new(big.Int).SetBytes(privKey.PrivateKey)
	if k.Cmp(curveOrder) >= 0 {
		return nil, "", errors.New("x509: invalid elliptic curve private key value")
	}
	return privKey.PrivateKey, curveName, nil
}
