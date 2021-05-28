package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	printList(l) //输出list的值,01234
	fmt.Println("")
	fmt.Println(l.Front().Value) //输出首部元素的值,0
	fmt.Println(l.Back().Value)  //输出尾部元素的值,4
	l.InsertAfter(6, l.Front())  //首部元素之后插入一个值为6的元素

	printList(l) //输出list的值,061234

	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换

	printList(l) //输出list的值,601234

	fmt.Println("")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部

	printList(l) //输出list的值,460123

	fmt.Println("")
	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾

	printList(l2) //输出l2的值,460123

	fmt.Println("")
	l.Init()           //清空l
	fmt.Print(l.Len()) //0

	printList(l) //输出list的值,无内容

}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}
