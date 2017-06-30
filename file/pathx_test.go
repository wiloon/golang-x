package file

import (
	"testing"
	"fmt"
)

func TestPathx(t *testing.T) {
	var str1, str2 string
	str1 = GetCurrentDirectory()
	fmt.Println("current:", str1)
	str2 = GetParentDirectory(str1)
	fmt.Println("parent:", str2)
}
