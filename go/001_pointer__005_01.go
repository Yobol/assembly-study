package main

import "unsafe"

func main() {
	println(size())
}

//go:noinline
func size() (o uintptr) {
	o = unsafe.Sizeof(o)
	return
}

// go build 001_pointer__005_01.go
// go tool objdump -S -s 'main.size' 001_pointer__005_01.exe

// 1.23.4 结果如下
// TEXT main.size(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__005_01.go
//         return
//   0x469a40              b808000000              MOVL $0x8, AX
//   0x469a45              c3                      RET

// unsafe 包并不是标准库的一部分，而是 Go 语言的一个特殊包，
// 本质上是一组 keyword，用于在编译时绕过 Go 语言的类型安全检查。
// 为什么这么设计？根本原因还是处于安全考虑。