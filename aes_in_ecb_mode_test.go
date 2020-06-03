package cryptopals

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDecryptAESinECBMode(t *testing.T){
	t.Run("correctly decrypts with valid key", func(t *testing.T){
		fn := "7.txt"
		key := "YELLOW SUBMARINE"

		data, err := ioutil.ReadFile(fn)
		if err != nil {
			t.Fatalf("Expected no error but got %v\n", err)
		}
		decrypted := DecryptAESInECBMode(DecodeBase64([]byte(data)), []byte(key))
		fmt.Printf("Decrypted text:\n%s\n", decrypted)
	})
}

