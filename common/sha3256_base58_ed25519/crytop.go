package sha3256_base58_ed25519

import (
	"hash"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
	"github.com/btcsuite/btcutil/base58"
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
	return "",""
}

func (c *Sha3256_base58_ed25519)Sign(priv string, msg string) string {
	return ""
}

func (c *Sha3256_base58_ed25519)Verify(pub string, msg string, sig string) bool {
	return true
}