package main

import (
	"fmt"
	"regexp"
)

var text = `20170215 15:30:20.075326 [127.0.0.1-thread-43054 ] INFO   server - key[key0] host[192.168.1.1:8000] request data:PREFIX,0,0,0,0,0,000 - abc.java::796`

func main() {

	test1()
}


func test1(){
	// 查找连续的小写字母
	//reg := regexp.MustCompile(`(d{6} dd:dd:dd.d{6}) .*(RG.*?) `)
	reg := regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}\.\d{6}).*(PREFIX.*?) .*`)

	result := reg.FindSubmatch([]byte(text))

	fmt.Printf("%s\n", result[0])
	fmt.Printf("%s\n", result[1])
	fmt.Printf("%s\n", result[2])
}


func  test0(){
	// 查找连续的小写字母
	reg := regexp.MustCompile(`PREFIX.*? `)
	//fmt.Printf("%q\n", reg.FindAllString(text, -1))
	fmt.Printf("%q\n", reg.FindString(text))
	// ["ello" "o"]
}


func test2() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	s := []byte(`abc-def-ghi abc+def+ghi`)

	// 查找所有匹配结果
	for _, one := range reg.FindAll(s, -1) {
		fmt.Printf("%s\n", one)
	}
}
