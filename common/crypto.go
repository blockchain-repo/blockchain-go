package common

import (
	"os/user"

	"unichain-go/common/sha3256_base58_ed25519"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var regStruct map[string]Crypto

type Crypto interface {
	//hash
	Hash(str string) string
	//encode
	Encode(b []byte) string
	Decode(str string) []byte
	//encrypt
	GenerateKeypair(seed ...string) (pub string, priv string)
	Sign(priv string, msg string) string
	Verify(pub string, msg string, sig string) bool
}

func init() {
	regStruct = make(map[string]Crypto)
	regStruct["sha3256/base58/ed25519"] = &sha3256_base58_ed25519.Sha3256_base58_ed25519{}
}

func GetCrypto() Crypto {
	var c Crypto
	_user, err := user.Current()
	if err != nil {
		logs.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	iniConfig, err := config.NewConfig("json", fileName)
	if err != nil {
		return nil
	}
	str := iniConfig.String("Crypto")
	c, ok := regStruct[str]
	if !ok {
		return nil
	}
	return c
}

func SetCrypto(cryptoType string) Crypto {
	c, ok := regStruct[cryptoType]
	if !ok {
		return nil
	}
	return c
}
