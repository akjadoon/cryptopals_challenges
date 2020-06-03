package cryptopals

import (
	"testing"
)

func TestDetectSingleByteXOR(t *testing.T){
	t.Run("runs", func(t* testing.T){
		fn := "4.txt"
		DetectSingleByteXOR(fn)
	})
}