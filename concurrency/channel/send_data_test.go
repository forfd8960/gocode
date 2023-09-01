package channel

import (
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDQueue_Enqueue(t *testing.T) {
	dq := &DQueue{
		queue: make(chan int, 1),
	}

	ch := make(chan struct{}, 2)

	go func() {
		var wg sync.WaitGroup
		for i := 1; i <= 10; i++ {
			v := i

			wg.Add(1)
			go func() {
				defer wg.Done()
				if dq.Enqueue(v) {
					log.Printf("enqueue data: %d\n", v)
				} else {
					log.Printf("enqueue data failed: %d\n", v)
				}
			}()
		}
		wg.Wait()

		ch <- struct{}{}
	}()

	go func() {
		count := 0
		for {
			if count == 10 {
				break
			}

			v, ok := dq.Dequeue()
			if ok {
				count += 1
				log.Printf("dequeue data: %d\n", v)
			} else {
				log.Printf("no data in queue\n")
			}
		}
		ch <- struct{}{}
	}()

	<-ch
	<-ch
}

func TestEnqueue(t *testing.T) {
	queue := make(chan int, 1)
	queue <- 1
	assert.False(t, Enqueue(queue, 10))
}
