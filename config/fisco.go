package config

import (
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
)

type GoSdk struct {
	Client   *client.Client                 `json:"client"`
	Contract map[string]*bind.BoundContract `json:"contract"`
}
