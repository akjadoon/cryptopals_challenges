package cryptopals

import (
	"bytes"
)
func DetectAESInECBMode(b []byte) (int, int) {
	bestLineNum := -1
	bestCount := 0

	for lineNum, line := range bytes.Split(b, []byte("\n")) {
		counts := map[string]bool{}
		repeats := 0
		line = DecodeBase64(line)
		for i := 0; i + 16 < len(line); i += 16 {
			block := string(line[i:i+16])
			if _, exists := counts[block]; exists {
				repeats += 1
			} else {
				counts[block] = true
			}
			if repeats > bestCount {
				bestCount = repeats
				bestLineNum = lineNum + 1
			}
		}
	}
	return bestLineNum, bestCount
}