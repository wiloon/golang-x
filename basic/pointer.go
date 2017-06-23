package basic

import (
	"fmt"
)

func PointerTest() {

	var i int  // i 的类型是int型
	i = 1      // i 的值为 1;
	var p *int // p 的类型是[int型的指针]
	p = &i     // p 的值为 [i的地址]

	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

	*p = 2 // *p 的值为 [[i的地址]的指针] (其实就是i嘛),这行代码也就等价于 i = 2
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

	i = 3 // 验证想法
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

}
