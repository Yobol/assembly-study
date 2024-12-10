package main

var n int

func main()  {
	println(addr())
}

//go:noinline
func addr() (p *int) {
	return &n
}

// go build 001_pointer__002_01.go
// go tool objdump -S -s 'main.addr' 001_pointer__002_01.exe

// 1.23.4 结果如下
// TEXT main.addr(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__002_01.go
//         return &n
//   0x469a40              488d0501d30c00          LEAQ main.n(SB), AX
//   0x469a47              c3                      RET
