package hashx

import (
	"testing"
	"fmt"
)

func TestHashStr(t *testing.T) {
	h := CalculateHash("a")
	fmt.Println("hash:", h)
	if h != 97 {
		t.Error("hash calculate failed.")
	}
}

func TestHashStr0(t *testing.T) {
	h := CalculateHash("test-string")
	fmt.Println("hash:", h)
}
