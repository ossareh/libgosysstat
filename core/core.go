package core

import (
	"log"
	"time"
)

type Stat struct {
	Type   string
	Values []int64
}

type ResultProcessor interface {
	Process() ([]Stat, error)
}

func StatProcessor(processor ResultProcessor, interval time.Duration, c chan *[]Stat) {
	tick := time.Tick(interval * time.Second)
	for {
		select {
		case <-tick:
			d, err := processor.Process()
			if err == nil {
				c <- &d
			} else {
				log.Println(err)
			}
		}
	}
}
