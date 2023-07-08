package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := "2147483649"
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("num: ", num)
}
