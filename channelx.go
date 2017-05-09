package main

func main() {
	// 创建一个channel用以同步goroutine
	done := make(chan bool)

	// 在goroutine中执行输出操作
	go func() {
		println("goroutine message")

		// 告诉main函数执行完毕.
		// 这个channel在goroutine中是可见的
		// 因为它是在相同的地址空间执行的.
		done <- true
	}()

	println("main function message")
	<-done // 等待goroutine结束
}