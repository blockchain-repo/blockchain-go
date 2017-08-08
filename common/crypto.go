package common

import (
	"unichain-go/common/sha3256_base58_ed25519"
)

var regStruct map[string]Crypto

type Crypto interface {
	//hash
	Hash(str string) string
	//encode
	Encode(b []byte) string
	Decode(str string) []byte
	//encrypt
	GenerateKeypair(seed ...string) (pub string,priv string)
	Sign(priv string, msg string) string
	Verify(pub string, msg string, sig string) bool
}


func init() {
	regStruct = make(map[string]Crypto)
	regStruct["sha3256/base58/ed25519"] = &sha3256_base58_ed25519.Sha3256_base58_ed25519{}
}


