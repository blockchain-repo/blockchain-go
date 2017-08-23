package common

import "fmt"

func ExampleCrypto() {
	c := GetCrypto()
	fmt.Println(c.Hash("jihao"))
	fmt.Println(c.GenerateKeypair("6hXsHQ4fdWQ9UY1XkBYCYRouAagRW8rXxYSLgpveQNYY"))
	msg := "hello unichain 2017"
	pub := "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	pub2 := "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	pri := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	sig := "48cpAsUuNf6qKCMFFKitSNjaA8nfPM4o7MacVp8U3QVMbVUr34SSRTTpahi3WEv3GaF2bVWG7J4SLTojgDoacLxT"
	sig2 := c.Sign(pri, msg)
	fmt.Println(sig, sig2)
	fmt.Println(c.Verify(pub, msg, sig))
	fmt.Println(c.Verify(pub2, msg, sig))

	// Output:
	//ced908fa6159f505e4d3f46e2c5a45f65476fe612856b5e347b6103be2c239f3
	//BbfY4Dc5s8dTP1Z1yixnetezRKYREHqwbP8GQGh3WyVS 6hXsHQ4fdWQ9UY1XkBYCYRouAagRW8rXxYSLgpveQNYY
	//48cpAsUuNf6qKCMFFKitSNjaA8nfPM4o7MacVp8U3QVMbVUr34SSRTTpahi3WEv3GaF2bVWG7J4SLTojgDoacLxT 48cpAsUuNf6qKCMFFKitSNjaA8nfPM4o7MacVp8U3QVMbVUr34SSRTTpahi3WEv3GaF2bVWG7J4SLTojgDoacLxT
	//true
	//false
}
