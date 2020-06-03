package cryptopals

import (
	"fmt"
	"crypto/aes"
)

func DecryptAESInECBMode(b []byte, key []byte) []byte {
	fmt.Printf("Length of ciphertext%v\n",  len(b) )
	output := make([]byte, len(b))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	for i := 0; i + 16 < len(b); i+=16 {
		cipher.Decrypt(output[i:i+16], b[i:i+16])
	}
	return output
}