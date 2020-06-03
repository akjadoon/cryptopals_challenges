package cryptopals

import (
    "testing"
)


func TestHexToBase64(t *testing.T) {
	t.Run("decodes ascii []byte to hex []byte", func(t *testing.T){
		// Converts UTF-8/Ascii string to byte-slice
		hexStr := []byte("012abc")
		expected := []byte{1, 42, 188}
		got, err := DecodeHex(hexStr)

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}

		if len(expected) != len(got) {
			t.Errorf("Expected byte slice of length %d but got %d", len(expected), len(got))
		}

		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("Expected %v at position %v but got %v", expected[i], i, got[i])
			}
		}
	})

	t.Run("encodes []byte to base64 []byte", func(t* testing.T){
		input := []byte{0, 2, 4}
		expected := []byte("AAIE")
		
		got := EncodeBase64(input)

		if len(expected) != len(got) {
			t.Fatalf("Expected byte slice of length %d but got %d", len(expected), len(got))
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("Expected %v at position %v but got %v", expected[i], i, got[i])
			}
		}
	})

	t.Run("converts hex encoded []byte to base64 encoded []byte", func(t* testing.T){
		input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
		expected := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
		got, err := HexToBase64(input)

		if err != nil {
			t.Fatalf("Expected no error but got %v", err)
		}

		for i := range expected {
			if got[i] != expected[i] {
				t.Errorf("Expected %v at position %v but got %v", expected[i], i, got[i])
			}
		}
	})
	
}

