package file

import (
	"testing"
	"log"
)

func TestReader(t *testing.T) {
	log.Println(Read("foo.txt"))
}
