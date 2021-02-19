package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
	fmt.Println("b return:", b()) // 打印结果为 b return: 2
	fmt.Println("c return:", c()) // 打印结果为 b return: 2
}

//匿名返回值的情况
func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

//有名返回值的情况
func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}

func c() (i int) {
	defer func() {
		if err := recover(); err != nil {

			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Printf("recover, %s\n", string(buf[:n]))
			i++
			fmt.Println("c defer2-1:", i)
		}

		fmt.Println("c defer2:", i) // 打印结果为 b defer2: 2
	}()
	panic(fmt.Sprintf("%v", i))
	return i // 或者直接 return 效果相同
}
