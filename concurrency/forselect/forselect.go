package forselect

import (
	"fmt"
	"time"

	"github.com/thoas/go-funk"
)

type Item struct {
	ID string
}

func IteratorItems(items []*Item, stop chan struct{}) {
	timer := time.NewTicker(1 * time.Second)
	defer timer.Stop()

	itemsChunks := funk.Chunk(items, 5)
	chunkItems := itemsChunks.([][]*Item)

	var count = 0
	for {
		if count >= len(chunkItems) {
			break
		}

		select {
		case <-timer.C:
			fmt.Printf("tm: %v, chunk: %v\n", time.Now(), chunkItems[count])
			count += 1
		case <-stop:
			fmt.Printf("stoped iter...")
			return
		}
	}
}
