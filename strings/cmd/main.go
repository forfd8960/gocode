package main

import (
	"fmt"

	"github.com/forfd8960/gocode/strings"
)

func main() {
	ss := strings.NewStrSplit("a b c", " ")

	iter := ss.Split()

	v := iter.Next()
	for v != "" {
		fmt.Println("val: ", v)
		v = iter.Next()
	}

	// fmt.Println("val1: ", iter.Next())
	// fmt.Println("val2: ", iter.Next())
	// fmt.Println("val3: ", iter.Next())
	// fmt.Println("val4: ", iter.Next())
}
