package basic

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}
func startProcess() {
	if proc, err := Start("ping", "-c 3", "www.baidu.com"); err == nil {
		proc.Wait()
	}
	if proc, err := Start("zsh"); err == nil {
		proc.Wait()
	}
}
func foo() {
	//在我们的例子中，我们将执行 ls 命令。Go 需要提供我们需要执行的可执行文件的绝对路径，所以我们将使用exec.LookPath 来得到它（大概是 /bin/ls）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	//Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}
	//Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量。
	env := os.Environ()
	//这里是 os.Exec 调用。如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func exec0() {
	cmd := exec.Command("ls") //查看当前目录下文件
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
