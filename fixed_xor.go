package cryptopals

import (
	"encoding/hex"
	"errors"
)

func FixedXOR(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		return nil
	}
	output := make([]byte, len(a))
	
	for i := range a {
		output[i] = a[i] ^ b[i]
	}
	return output
}

func FixedXORHex (a []byte, b []byte) ([]byte, error) {
	aBytes, err := DecodeHex(a)
	if err != nil {
		return nil, err
	}
	bBytes, err := DecodeHex(b)
	if err != nil {
		return nil, err
	}

	if len(a) != len(b) {
		return nil, errors.New("Inputs a and b not equal in length")
	}

	xorBytes := FixedXOR(aBytes, bBytes)
	output := make([]byte, hex.EncodedLen(len(xorBytes)))
	hex.Encode(output, xorBytes)

	return output, nil
}