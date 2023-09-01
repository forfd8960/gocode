package channel

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type data struct {
	mu     sync.Mutex
	dataCh chan int
	stop   int32
}

func NewData() *data {
	d := &data{
		dataCh: make(chan int, 1),
	}

	go d.ticker()

	return d
}

func (d *data) ticker() {
	for !d.isStop() {
		d.waitForData()
	}
}

func (d *data) isStop() bool {
	return atomic.LoadInt32(&d.stop) == 1
}

func (d *data) waitForData() {
	wait := 100 + (rand.Int63() % 30)
	duration := time.Duration(wait) * time.Millisecond
	time.Sleep(duration)

	log.Printf("after wait: %s\n", duration.String())

	if d.receviedData() {
		log.Println("received data")
	}
}

func (d *data) SendData(val int) {
	go func() {
		d.mu.Lock()
		defer d.mu.Unlock()
		if len(d.dataCh) >= 1 {
			log.Printf("dataCh is full for: %d\n", val)
			return
		}

		log.Printf("enqueue %d to channel", val)
		d.dataCh <- val
	}()
}

func (d *data) receviedData() bool {
	select {
	case dt := <-d.dataCh:
		log.Printf("received data: %d", dt)
		return true
	default:
		return false
	}
}
