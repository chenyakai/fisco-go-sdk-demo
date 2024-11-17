package contract

import "math/big"

type Output struct {
	Id         *big.Int
	UserID     string
	UserPK     string
	Nounce     string
	Message    string
	Cipher     string
	Attachment string
	Signtime   *big.Int
}
