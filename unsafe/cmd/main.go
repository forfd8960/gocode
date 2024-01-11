package main

import (
	"fmt"
	"unsafe"
)

func main() {
	Offsetof()
}

type X struct {
	a int64
	b bool
	c string
}

func Offsetof() {
	var sizeOfX = unsafe.Sizeof(X{})
	fmt.Println("x size: ", sizeOfX) // 32

	var x = X{}
	/*
		offset of: a in X: 0
		offset of: a in X: 8
		offset of: a in X: 16
	*/
	fmt.Printf("offset of: a in X: %d\n", unsafe.Offsetof(x.a))
	fmt.Printf("offset of: a in X: %d\n", unsafe.Offsetof(x.b))
	fmt.Printf("offset of: a in X: %d\n", unsafe.Offsetof(x.c))
}
