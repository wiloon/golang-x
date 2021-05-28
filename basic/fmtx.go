package basic

import (
	"flag"
	"fmt"
)

func Fmtx() {
	qps := flag.Int("qps", 1, "qps")

	fmt.Printf("foo:%s\n", "bar")

	fmt.Printf("foo:%v\n", *qps)
}
