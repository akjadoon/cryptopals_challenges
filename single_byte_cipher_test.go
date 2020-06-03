package cryptopals

import (
	"testing"
	"fmt"
)

func TestSingleByteCipher(t *testing.T) {
	t.Run("decodes single byte cipher", func(t *testing.T){
		input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
		s, _, _ := DecodeSingleByteCipher(input)
		fmt.Printf("Deciphered to \n%v\n", string(s))
	})
}