package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
	"math/bits"
)
func HammingDistance(a []byte, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Expected a and b to have same length.") 
	}
	distance := 0

	for i := range a {
		distance += bits.OnesCount8(a[i] ^ b[i])
	}
	return distance, nil
}


func ShortestHammingDistanceKeyLength(b []byte, start int, end int) int {
	bestKey := 0
	shortestDistance := math.MaxFloat64
	for keylen := start; keylen <= end; keylen++ {
		comparisons := 0
		totalDistance := 0.0
		for i:= 0; i + (2 * keylen) < len(b); i += (2 * keylen) {
			dist, err := HammingDistance(b[i:i+keylen], b[i+keylen:i+(2*keylen)])
			if err != nil {
				panic(err)
			}
			totalDistance += float64(dist) / float64(keylen)
			comparisons += 1
		}
		if float64(totalDistance) / float64(comparisons) < float64(shortestDistance) {
			bestKey = keylen
			shortestDistance = float64(totalDistance) / float64(comparisons)
		}
	}
	return bestKey
}

func DecodeBase64(b []byte) []byte {
	output := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	base64.StdEncoding.Decode(output, b)
	return output
}

func DecryptRepeatingKeyXOR(b []byte, keylen int) ([]byte, []byte) {
	blocks := make([][]byte, keylen)
	for i:= 0; i < keylen; i++ {
		if i < (len(b) % keylen) {
			blocks[i] = make([]byte, (len(b) / keylen) + 1)
		} else {
			blocks[i] = make([]byte, len(b) / keylen)
		}
		for j:= 0; j < len(blocks[i]); j++ {
			blocks[i][j] = b[i + (j * keylen)]
		}
	}
	output := make([]byte, len(b))
	key := make([]byte, keylen)
	for i := range blocks {
		hexEncoded := make([]byte, hex.EncodedLen(len(blocks[i])))
		hex.Encode(hexEncoded, blocks[i])
		deciphered, _, keyByte := DecodeSingleByteCipher(hexEncoded)
		key[i] = keyByte 
		for j := range deciphered {
			output[i + (j * keylen)] = deciphered[j]
		}
	}

	return output, key
}