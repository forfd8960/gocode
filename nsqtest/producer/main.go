package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic(err)
	}
	err = p.Publish("NSQ_Topic", []byte("Hello NSQ, Hello Hello Hello"))
	if err != nil {
		log.Panic(err)
	}
}
