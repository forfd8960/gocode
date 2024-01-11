package unsafe

import (
	"fmt"
	"unsafe"
)

type X struct {
	a int64
	b bool
	c string
}

func Offsetof() {
	var sizeOfX = unsafe.Sizeof(X{})
	fmt.Println("x size: ", sizeOfX)
}
