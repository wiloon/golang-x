package hashx

import (
	"fmt"
	"math"
	"flag"
)

func main() {
	k := flag.String("k", "", "hash key string")
	flag.Parse()

	fmt.Println("key:" + *k)

	var hashValue = CalculateHash(*k)
	fmt.Println("hash:", hashValue)
}

func CalculateHash(key string) float64 {
	var hash int32 = 0;
	for _, c := range key {
		hash = 31*hash + int32(c)
	}
	return math.Abs(float64(hash % 1024))
}
