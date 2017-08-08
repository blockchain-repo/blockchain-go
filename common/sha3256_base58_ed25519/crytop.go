package sha3256_base58_ed25519

type Sha3256_base58_ed25519 struct {

}

//hash
func (c *Sha3256_base58_ed25519)Hash(str string) string {
	return ""
}

//encode
func (c *Sha3256_base58_ed25519)Encode(b []byte) string {
	return ""
}
func (c *Sha3256_base58_ed25519)Decode(str string) []byte {
	return []byte("")
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