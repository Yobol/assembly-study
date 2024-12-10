package main

import "unsafe"

func main()  {
	n := 2
	convert(&n)
	println(n)
}

//go:noinline
func convert(p *int) {
	q := (*int32)(unsafe.Pointer(p))
	*q = 0
}

// go build 001_pointer__003_01.go
// go tool objdump -S -s 'main.addr' 001_pointer__003_01.exe

// 1.23.4 结果如下
// TEXT main.convert(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__003_01.go
//         *q = 0
//   0x469a60              c70000000000            MOVL $0x0, 0(AX)
// }
//   0x469a66              c3                      RET
