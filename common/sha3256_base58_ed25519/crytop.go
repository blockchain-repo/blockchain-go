package sha3256_base58_ed25519

import (
	"hash"
	"encoding/hex"

	"unichain-go/log"

	"golang.org/x/crypto/sha3"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
	"bytes"
)

type Sha3256_base58_ed25519 struct {

}

//hash
func (c *Sha3256_base58_ed25519)Hash(str string) string {
	var hash hash.Hash
	var x string = ""
	hash = sha3.New256()
	if hash != nil {
		hash.Write([]byte(str))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x

}

//encode
func (c *Sha3256_base58_ed25519)Encode(b []byte) string {
	return base58.Encode(b)
}
func (c *Sha3256_base58_ed25519)Decode(str string) []byte {
	return base58.Decode(str)
}

//encrypt
func (c *Sha3256_base58_ed25519)GenerateKeypair(seed ...string) (pub string,priv string) {
	var publicKeyBytes, privateKeyBytes []byte
	var err error
	if len(seed) >= 1 {
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(bytes.NewReader(base58.Decode(seed[0])))
	} else  {
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(nil)
	}
	if err != nil {
		log.Error(err.Error())
	}
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes[0:32])
	return publicKeyBase58, privateKeyBase58
}

func (c *Sha3256_base58_ed25519)Sign(priv string, msg string) string {
	return ""
}

func (c *Sha3256_base58_ed25519)Verify(pub string, msg string, sig string) bool {
	return true
}