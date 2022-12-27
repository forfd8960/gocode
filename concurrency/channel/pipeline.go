package channel

import "fmt"

// Pipe chain of the channels
func Pipe(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("recevied: %d\n", n)
			out <- n * n
		}
		close(out)
	}()

	return out
}
