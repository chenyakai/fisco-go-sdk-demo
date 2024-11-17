package utils

import (
	goecdsa "crypto/ecdsa"
	"errors"
	"fisco-go-sdk-demo/global"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

/*
公用发送交易组装器
*/
func SendTransaction(name, method string, params ...interface{}) any {
	_, receipt, ok := global.GoSdk.Contract[name].Transact(global.GoSdk.Client.GetTransactOpts(), method, params...)

	if ok != nil {
		fmt.Println("txError=>", ok)
		return nil
	}
	json, wrong := abi.JSON(strings.NewReader(global.Config.Contract[name].Abi))
	if wrong != nil {
		fmt.Println("wrong==>", wrong)
	}

	var (
		result = new(any)
	)
	//合约方法名
	task := json.Unpack(&result, method, common.FromHex(receipt.Output))
	if task != nil {
		fmt.Println("task==>", task)
	}
	return *result
}

// SendCall 获取链上信息
func SendCall(name, method string, out interface{}, params ...interface{}) any {
	err := global.GoSdk.Contract[name].Call(global.GoSdk.Client.GetCallOpts(), out, method, params...)
	if err != nil {
		fmt.Println("SendCall err==>", err)
	}
	return out
}

func NewKeyedTransactor(key *goecdsa.PrivateKey) *bind.TransactOpts {
	//key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
}

// SendTransactionByKey 使用私钥发送交易
func SendTransactionByKey(name, method string, privateKey *goecdsa.PrivateKey, params ...interface{}) any {
	_, receipt, ok := global.GoSdk.Contract[name].Transact(NewKeyedTransactor(privateKey), method, params...)

	if ok != nil {
		fmt.Println("txError=>", ok)
		return nil
	}
	json, wrong := abi.JSON(strings.NewReader(global.Config.Contract[name].Abi))
	if wrong != nil {
		fmt.Println("wrong==>", wrong)
	}

	var (
		result = new(any)
	)
	//合约方法名
	task := json.Unpack(&result, method, common.FromHex(receipt.Output))
	if task != nil {
		fmt.Println("task==>", task)
	}
	return *result
}

// SendCallByKey SendCall 获取链上信息
func SendCallByKey(name, method string, privateKey *goecdsa.PrivateKey, out interface{}, params ...interface{}) any {
	//clientAuth := global.GoSdk.Client.GetTransactOpts
	//global.GoSdk.Client.SetTransactOpts(NewKeyedTransactor(privateKey))
	//clientAuth.GasLimit = big.NewInt(30000000)
	clientCallOpts := &bind.CallOpts{From: NewKeyedTransactor(privateKey).From}
	err := global.GoSdk.Contract[name].Call(clientCallOpts, out, method, params...)
	if err != nil {
		fmt.Println("SendCall err==>", err)
	}
	return out
}
