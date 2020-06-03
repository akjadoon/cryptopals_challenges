package cryptopals

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func DetectSingleByteXOR(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	lines := strings.Split(string(dat), "\n")
	maxScore := 0.0
	maxScoreStr := ""
	for _, l := range lines {
		check(err)
		decoded, score, _ := DecodeSingleByteCipher([]byte(l))
		if score > maxScore {
			maxScore = score
			maxScoreStr = string(decoded)
		}
		// fmt.Printf("%s\n", decoded)
	}
	check(err)
	fmt.Printf("%v\n", maxScoreStr)
	return nil
}

func StripChar(b []byte, char byte) []byte {
	instances := 0
	for i := range b {
		if b[i] == char {
			instances += 1
		}
	}
	output := make([]byte, len(b) - instances)
	cur := 0
	for i := range b {
		if b[i] != char {
			output[cur] = b[i]
			cur += 1
		}
	}
	return output
}