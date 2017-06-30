package tcp

import (
	"bytes"
	"fmt"
	"bufio"
	"encoding/binary"
	"io"
)

func ScannerTest() {
	// 创建一个包，版本是V1，数据是ABCDEFGHIJK，大小是11
	var pkg Package
	pkg.Version[0] = 'V'
	pkg.Version[1] = 1
	pkg.Data = []byte("ABCDEFGHIJK")
	pkg.Datalen = int16(len(pkg.Data))
	fmt.Println(pkg)
	fmt.Println(&pkg)

	// 打包成二进制数据
	var buf bytes.Buffer
	pkg.Pack(&buf)

	// 从二进制数据里面获取数据
	var pkg1 Package
	pkg1.Unpack(&buf)
	fmt.Println(&pkg1)
	// 模拟数据流，打包三个数据包
	pkg.Pack(&buf)
	pkg.Pack(&buf)
	pkg.Pack(&buf)

	// 创建Scanner，分析buf数据流(r io.Reader，换成net.Conn对象就是处理tcp数据流，自己连数据都不需要去收取)
	scanner := bufio.NewScanner(&buf)

	// 数据的分离规则，根据协议自定义
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'V' {
			if len(data) > 4 {
				var dataLen int16
				binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &dataLen)
				if int(dataLen)+4 <= len(data) {
					return int(dataLen) + 4, data[:int(dataLen)+4], nil
				}
			}
		}
		return
	}

	// 设置分离函数
	scanner.Split(split)

	// 获取分离出来的数据
	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

}

// 自定义协议的组包和拆包
type Package struct {
	Version [2]int8
	Datalen int16
	Data    []byte
}

func (p *Package) String() string {
	return fmt.Sprintf("Version:%d DataLen:%d Data:%s",
		p.Version, p.Datalen, p.Data)
}

func (p *Package) Pack(w io.Writer) {
	binary.Write(w, binary.BigEndian, p.Version)
	binary.Write(w, binary.BigEndian, p.Datalen)
	binary.Write(w, binary.BigEndian, p.Data)
}

func (p *Package) Unpack(r io.Reader) {
	binary.Read(r, binary.BigEndian, &p.Version)
	binary.Read(r, binary.BigEndian, &p.Datalen)
	if p.Datalen > 0 {
		p.Data = make([]byte, p.Datalen)
	}
	binary.Read(r, binary.BigEndian, &p.Data)
}
