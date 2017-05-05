package copy

import (
	"testing"

	"fmt"
)

func TestCopyFile(t *testing.T) {
	//src := "/home/wiloon/tmp/PIC_19700117_163128_C96.mp4"
	//dst := "/home/wiloon/tmp/out.mp4"

	//src:="c:/workspace/tmp/foo.txt"
	src:="//192.168.3.1/share/tmp/remote.txt"
	//src:="z:\\tmp\\remote.txt"
	dst:="c:/workspace/tmp/bar.txt"
	FileCopy(dst, src) // os.Args[1]为目标文件，os.Args[2]为源文件
	fmt.Println("done")
}
