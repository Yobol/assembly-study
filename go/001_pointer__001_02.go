package main

func main()  {
	n := int32(10)
	println(read32(&n))
}

//go:noinline
func read32(p *int32) (v int32) {
	v = *p
	return
}

// go build 001_pointer__001_02.go
// go tool objdump -S -s 'main.read' 001_pointer__001_02.exe

// 1.16 结果如下
// TEXT main.read32(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__001_02.go
//         v = *p
//   0x4626e0              488b442408              MOVQ 0x8(SP), AX
//   0x4626e5              8b00                    MOVL 0(AX), AX
//         return
//   0x4626e7              89442410                MOVL AX, 0x10(SP)
//   0x4626eb              c3                      RET

// 1.17 及以后版本结果如下，我这里用的 1.23.4
// TEXT main.read32(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__001_02.go
//         v = *p
//   0x469a60              8b00                    MOVL 0(AX), AX
//         return
//   0x469a62              c3                      RET

// 和 001_pointer__001_01.go 相比，该文件将 read() 函数修改为 read32() 函数
// 可以看到，用于复制指针存储地址的指令没有发生变化，内存寻址单元 0(AX) 也没有变化，
// 但是 MOVQ 变成了 MOVL，说明指针元素类型的变化会影响汇编的结果。
