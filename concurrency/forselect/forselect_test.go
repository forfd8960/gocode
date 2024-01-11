package forselect

import (
	"fmt"
	"testing"
	"time"
)

func TestIteratorItems(t *testing.T) {
	items := []*Item{}
	for i := 0; i < 50; i++ {
		items = append(items, &Item{ID: fmt.Sprintf("item%d", i)})
	}

	stop := make(chan struct{}, 1)
	go func() {
		IteratorItems(items, stop)
	}()
	time.Sleep(20 * time.Second)
	stop <- struct{}{}
}
