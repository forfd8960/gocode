package channel

import (
	"fmt"
	"log"
)

type DQueue struct {
	queue chan int
}

func (q *DQueue) Enqueue(d int) bool {
	select {
	case q.queue <- d:
		log.Printf("enqueue data success: %d\n", d)
		return true
	default:
		log.Printf("enqueue data failed: ")
		return false
	}
}

func (q *DQueue) Dequeue() (int, bool) {
	select {
	case v := <-q.queue:
		log.Printf("dequeue data success: %d\n", v)
		return v, true
	default:
		log.Printf("no data to dequeue\n")
		return -1, false
	}
}

func Enqueue(queue chan int, v int) bool {
	select {
	case queue <- v:
		fmt.Printf("send v to queue: %d\n", v)
		return true
	default:
		fmt.Printf("queue is full\n")
		return false
	}
}
