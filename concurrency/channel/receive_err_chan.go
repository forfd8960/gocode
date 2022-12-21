package channel

import (
	"fmt"
	"sync"
)

func ReceiveErrFromChan(num int, even bool) string {
	var wg sync.WaitGroup
	strChan := make(chan string, num)
	for i := 0; i < num; i++ {
		wg.Add(1)

		n := i
		go func() {
			defer wg.Done()
			fmt.Printf("running: %d\n", n)
			if even && n%2 == 0 {
				strChan <- "n is even"
			}

		}()
	}
	wg.Wait()
	close(strChan)

	if s := <-strChan; s != "" {
		fmt.Println(s)
		return s
	}

	return ""
}
