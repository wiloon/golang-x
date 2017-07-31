package basic

import (
	"fmt"
	"flag"
)

func Fmtx(){
	qps := flag.Int("qps", 1, "qps")

	fmt.Printf("foo:%s\n","bar");

	fmt.Printf("foo:%s\n",*qps);
}