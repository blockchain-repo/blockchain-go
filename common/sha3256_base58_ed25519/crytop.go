package sha3256_base58_ed25519

import (
	"encoding/hex"
	"hash"

	"unichain-go/log"

	"bytes"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"
)

type Sha3256_base58_ed25519 struct {
}

//hash
func (c *Sha3256_base58_ed25519) Hash(str string) string {
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
func (c *Sha3256_base58_ed25519) Encode(b []byte) string {
	return base58.Encode(b)
}
func (c *Sha3256_base58_ed25519) Decode(str string) []byte {
	return base58.Decode(str)
}

//encrypt
func (c *Sha3256_base58_ed25519) GenerateKeypair(seed ...string) (pub string, priv string) {
	var publicKeyBytes, privateKeyBytes []byte
	var err error
	if len(seed) >= 1 {
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(bytes.NewReader(c.Decode(seed[0])))
	} else {
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(nil)
	}
	if err != nil {
		log.Error(err.Error())
	}
	publicKeyBase58 := c.Encode(publicKeyBytes)
	privateKeyBase58 := c.Encode(privateKeyBytes[0:32])
	return publicKeyBase58, privateKeyBase58
}

func (c *Sha3256_base58_ed25519) Sign(priv string, msg string) string {
	pub, _ := c.GenerateKeypair(priv)
	privByte := base58.Decode(priv)
	pubByte := base58.Decode(pub)
	privateKey := make([]byte, 64)
	copy(privateKey[:32], privByte)
	copy(privateKey[32:], pubByte)
	sigByte := ed25519.Sign(privateKey, []byte(msg))
	return base58.Encode(sigByte)
}

func (c *Sha3256_base58_ed25519) Verify(pub string, msg string, sig string) bool {
	pubByte := c.Decode(pub)
	publicKey := make([]byte, 32)
	copy(publicKey, pubByte)
	return ed25519.Verify(publicKey, []byte(msg), c.Decode(sig))
}
