package cryptopals

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestHammingDistance(t *testing.T){
	t.Run("correctly calculates hamming distance b/w 2 []byte", func(t *testing.T){
		a := []byte("this is a test")
		b := []byte("wokka wokka!!!")
		expected := 37
		dist, err := HammingDistance(a, b)
		if err != nil {
			t.Fatalf("Expected no error but got %v\n", err)
		}
		if dist != expected {
			t.Errorf("Expected distance %v but got %v", expected, dist)
		}
	})
}

func TestShortestHammingDistanceKeyLength(t *testing.T){
	t.Run("correctly finds the key length with the shortest hamming distance for []byte", func(t *testing.T){
		dat, err := ioutil.ReadFile("6.txt")
		if err != nil {
			t.Fatalf("Expected no error but got %v\n", err)
		}
		dist := ShortestHammingDistanceKeyLength(DecodeBase64(dat), 2, 40)

		fmt.Printf("Shortest key has length %d\n", dist)
	})
}

func TestBase64Decode(t *testing.T){
	t.Run("correctly decodes base64 encoded []byte", func(t *testing.T){
		input := []byte("aGVsbG8gd29ybGQh")
		expected := "hello world!"
		got := DecodeBase64(input)
		if string(got) != expected {
			t.Errorf("Expected %s in decoded string but got %s", expected, got)
		}
	})
}

func TestDecryptRepeatingKeyXOR(t *testing.T){
	t.Run("correctly decrypts text", func(t *testing.T){
		fn := "6.txt"
		dat, err := ioutil.ReadFile(fn)
		if err != nil {
			t.Fatalf("Expected no error but got %v\n", err)
		}
		output, key := DecryptRepeatingKeyXOR(DecodeBase64(dat), 29)
		fmt.Printf("Decrypted Output\n%s\n", output)
		fmt.Printf("Key Used\n%s\n", key)
	})
}