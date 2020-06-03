package cryptopals

import (
	"testing"
)

func TestPKCS7Padding(t *testing.T){
	t.Run("correctly pads []byte to 20 bytes", func(t *testing.T){
		length := 20 
		input := []byte("YELLOW SUBMARINE")
		expected := "YELLOW SUBMARINE\x04\x04\x04\x04"
		got := PKCS7Pad(input, length)
		if string(got) != expected {
			t.Errorf("Expected %s, but got %s", expected, got)
		}
	})
}