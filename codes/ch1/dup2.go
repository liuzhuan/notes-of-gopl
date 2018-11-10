package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// func Open(name string) (*File, error)
			f, err := os.Open(arg)
			if err != nil {
				// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
				// Who is io.Writer?
				// type Writer interface { Write(p []byte) (n int, err error) }
				// Who is os.Stderr?
				// Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
				// What does NewFile do?
				// func NewFile(fd uintptr, name string) *File
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// os.File 是一个结构体，有 Read 方法，满足 io.Reader 接口
func countLines(f *os.File, counts map[string]int) {
	// bufio.NewScanner 期望的参数格式是 io.Reader
	// io.Reader 是接口，其定义如下：
	// type Reader interface {
	//     Read(p []byte) (n int, err error)
	// }

	// bufio.NewScanner 返回的类型是 Scanner，这是一个结构体
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}