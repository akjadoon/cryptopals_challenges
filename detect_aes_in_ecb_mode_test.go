package cryptopals

import (
	"io/ioutil"
	"fmt"
	"testing"
)

func TestDetectAESECBMode(t *testing.T){
	t.Run("correctly detects string encrypted in ECB", func(t *testing.T){
		fn := "8.txt"

		data, err := ioutil.ReadFile(fn)
		if err != nil {
			t.Fatalf("Expected no error but got %v\n", err)
		}
		lineNum, count := DetectAESInECBMode(data)
		fmt.Printf("Repeat block count: %v\n line:\n%v\n", count, lineNum)
	})
}