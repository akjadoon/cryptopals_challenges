package cryptopals

import (
	"encoding/hex"
	"encoding/base64"
)

func DecodeHex(input []byte) ([]byte, error) {
	output := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(output, input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func EncodeBase64(input []byte) ([]byte) {
	output := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(output, input)
	return output
}

func HexToBase64(input []byte) ([]byte, error) {
	bytes, err := DecodeHex(input)
	if err != nil {
		return nil, err
	}
	return EncodeBase64(bytes), nil
}