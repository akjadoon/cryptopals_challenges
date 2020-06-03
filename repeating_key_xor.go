package cryptopals

import (
	"encoding/hex"
)


func RepeatingKeyXOR(plaintext []byte, key []byte) []byte {
	rex := make([]byte, len(plaintext))
	for i := range plaintext {
		rex[i] = plaintext[i] ^ key[i % len(key)]
	}
	output := make([]byte, hex.EncodedLen(len(rex)))
	hex.Encode(output, rex)
	return output
}