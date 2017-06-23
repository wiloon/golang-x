package basic

func DeferTest0() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func DeferTest0F() (result int) {
	result = 0 //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	func() { //defer被插入到return之前执行，也就是赋返回值和ret指令之间
		result++
	}()
	return
}

func DeferTest1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func DeferTest1F() (r int) {
	t := 5
	r = t //赋值指令
	func() { //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		t = t + 5
	}()
	return //空的return指令
}

func DeferTest2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func DeferTest2F() (r int) {
	r = 1  //给返回值赋值
	func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return        //空的return
}
