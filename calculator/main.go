package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/forfd8960/gocode/calculator/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Expression >>")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		result, err := eval.Eval(text)
		fmt.Println("result: ", result)
		fmt.Println("err: ", err)
	}
}
