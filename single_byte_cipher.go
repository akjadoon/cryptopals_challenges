package cryptopals

import (
	"strings"
)

func DecodeSingleByteCipher(input []byte) ([]byte, float64, byte) {
	input, _ = DecodeHex(input)

	bestCount := 0.0
	bestByte := byte(0)

	for i:= 0; i < 255; i++ {
		score := 0.0
		// fmt.Printf("%v:\n", string([]byte{uint8(i)}))
		for j := range input {
			char := input[j] ^ byte(i)
			// fmt.Printf("%v", string([]byte{uint8(char)}))
			score += CharFrequencyScore(char)
			if score > bestCount {
				bestCount = score
				bestByte = byte(i)
			}
		}
		
	}

	deciphered := make([]byte, len(input))
	for i := range input {
		deciphered[i] = input[i] ^ bestByte
	}
	return deciphered, bestCount, bestByte
}

func CharFrequencyScore (b byte) float64 {
	freqs := map[string]float64 {
		"e": 12.02,
		"t": 9.1,
		"a": 8.12,
		"o": 7.68,
		"i": 7.31,
		"n": 6.95,
		" ": 6.60,
		"s": 6.28,
		"r": 6.02,
		"h": 5.92,
		"d": 4.32,
		"l": 3.98,
		"u": 2.88,
		"c": 2.71,
		"m": 2.61,
		"f": 2.30,
		"y": 2.11,
		"w": 2.09,
		"g": 2.03,
		"p": 1.82,
		"b": 1.49,
		"v": 1.11,
		"k": 0.69,
		"x": 0.17,
		"q": 0.11,
		"j": 0.10,
		"z": 0.07,
	}
	s := strings.ToLower(string(b))
	if score, exists := freqs[s]; exists {
		return score
	}
	return 0.0
}