package channel

import (
	"fmt"
	"testing"
)

func TestPipe(t *testing.T) {
	in := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	out := Pipe(in)
	for n := range out {
		fmt.Println(n)
	}
}
