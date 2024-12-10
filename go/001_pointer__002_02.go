package main

func main()  {
	println(addr())
}

//go:noinline
func addr() (p *int) {
	var n int
	return &n
}

// go build 001_pointer__002_02.go
// go tool objdump -S -s 'main.addr' 001_pointer__002_02.exe

// 1.23.4 结果如下
// TEXT main.addr(SB) D:/workspace/github.com/yobol/assembly-study/go/001_pointer__002_02.go
// func addr() (p *int) {
// 	0x469a40              493b6610                CMPQ SP, 0x10(R14)
// 	0x469a44              761b                    JBE 0x469a61
// 	0x469a46              55                      PUSHQ BP
// 	0x469a47              4889e5                  MOVQ SP, BP
// 	0x469a4a              4883ec10                SUBQ $0x10, SP
// 		  var n int
// 	0x469a4e              488d05cb5d0000          LEAQ type:*+22560(SB), AX
// 	0x469a55              e8e61cfaff              CALL runtime.newobject(SB)
// 		  return &n
// 	0x469a5a              4883c410                ADDQ $0x10, SP
// 	0x469a5e              5d                      POPQ BP
// 	0x469a5f              90                      NOPL
// 	0x469a60              c3                      RET
//   func addr() (p *int) {
// 	0x469a61              e89a88ffff              CALL runtime.morestack_noctxt.abi0(SB)
// 	0x469a66              ebd8                    JMP main.addr(SB)
