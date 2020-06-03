package cryptopals

import (
	"testing"
)

func TestFixedXOR(t *testing.T) {
	t.Run("correctly XORs []byte a with []byte b", func(t *testing.T){
		a := []byte{100, 150, 200}
		b := []byte{10, 40, 255}
		expected := []byte{110, 190, 55}
		got := FixedXOR(a, b)

		for i := range expected {
			if expected[i] != got[i] {
				t.Errorf("Expected %v at position %v but got %v", expected[i], i, got[i])
			}
		}
	})
	t.Run("correctly XORs hex encoded []byte", func(t* testing.T){
		a := []byte("1c0111001f010100061a024b53535009181c")
		b := []byte("686974207468652062756c6c277320657965")
		expected := []byte("746865206b696420646f6e277420706c6179")
		got, err := FixedXORHex(a, b)

		if err != nil {
			t.Fatalf("Expected no error but got %v", err)
		}

		for i := range expected {
			if expected[i] != got[i] {
				t.Errorf("Expected %v at position %v but got %v", expected[i], i, got[i])
			}
		}
	})
}