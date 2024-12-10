package main

import "unsafe"

func main() {
	s := []byte{'H', 'e', 'l', 'l', 'o'}
	str := convert(s)
	println(str)
	s[0] = 'h'
	println(str)
}

func convert(s []byte) string {
	return *((*string)(unsafe.Pointer(&s)))
}

// go run 001_pointer__004_01.go

// 输出：
// Hello
// hello

// 解释：
// 使用了 unsafe.Pointer 操作并不会创建新的字符串，而是直接将切片的底层数组指针转换为字符串指针。
// 因此，对切片的修改会影响到字符串的内容。