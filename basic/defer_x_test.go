package basic

import (
	"testing"
	"fmt"
)

func TestDefer(t *testing.T) {
	fmt.Println(DeferTest0())
	fmt.Println(DeferTest0F())

	fmt.Println(DeferTest1())
	fmt.Println(DeferTest1F())

	fmt.Println(DeferTest2())
	fmt.Println(DeferTest2F())
}

