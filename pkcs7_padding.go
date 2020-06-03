package  cryptopals

func PKCS7Pad(input []byte, length int) []byte {
	paddingLen := length - (len(input) % length)
	output := make([]byte, len(input) + paddingLen)
	copy(output, input)
	for i := len(input); i < len(input) + paddingLen; i++ {
		output[i] = byte(paddingLen)
	}
	return output
}