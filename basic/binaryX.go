package basic

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func binaryX() {
	int16buf := new(bytes.Buffer)
	i := uint16(172)
	binary.Write(int16buf, binary.LittleEndian, i)
	fmt.Println("write buf is:", int16buf.Bytes())

	var int16buf2 [2]byte
	binary.LittleEndian.PutUint16(int16buf2[:], uint16(172))
	fmt.Println("put buffer is :", int16buf2[:])

	ii := binary.LittleEndian.Uint16(int16buf2[:])
	fmt.Println("Get buf is :", ii)





}


func Ascii(){
	var c rune='a'
	var i int =98
	i1:=int(c)
	fmt.Println("'a' convert to",i1)
	c1:=rune(i)
	fmt.Println("98 convert to",string(c1))

	//string to rune
	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}
}