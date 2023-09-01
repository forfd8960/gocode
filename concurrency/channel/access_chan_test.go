package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewData(t *testing.T) {
	data := NewData()
	// rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 500; i++ {
			v := i
			ms := 10
			du := time.Duration(ms) * time.Millisecond
			time.Sleep(du)

			fmt.Printf("send data: %d after wait: %s\n", v, du.String())
			data.SendData(v)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 501; i <= 1000; i++ {
			v := i
			ms := 10
			du := time.Duration(ms) * time.Millisecond
			time.Sleep(du)

			fmt.Printf("send data: %d after wait: %s\n", v, du.String())
			data.SendData(v)
		}
	}()
	wg.Wait()
}
