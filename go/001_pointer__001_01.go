package main

func main()  {
	n := 10
	println(read(&n))
}

// go:noline 是一个编译器指令，用于告诉编译器不要内联该函数。注意，go:noline 和 // 之间不要有空格！
//go:noinline
func read(p *int) (v int) {
	v = *p
	return
}

// go build 001_pointer__001_01.go
// go tool objdump -S -s 'main.read' 001_pointer__001_01.exe

// 1.16 结果如下
// TEXT main.read(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__001_01.go
//         v = *p
//   0x4626e0              488b442408              MOVQ 0x8(SP), AX
//   0x4626e5              488b00                  MOVQ 0(AX), AX
//         return
//   0x4626e8              4889442410              MOVQ AX, 0x10(SP)
//   0x4626ed              c3                      RET

// 1.17 及以后版本结果如下，我这里用的 1.23.4
// TEXT main.read(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__001_01.go
//         v = *p
//   0x469a60              488b00                  MOVQ 0(AX), AX
//         return
//   0x469a63              c3                      RET

// 可以看到，1.17 及以后版本函数入参和返回值在寄存器中，而不是栈中，整体性能提升 5% 左右。
