package core

import (
	"log"
	"time"
)

type Stat interface {
	Type() string
	Values() []uint64
}

type ResultProcessor interface {
	Process() ([]Stat, error)
}

func StatProcessor(processor ResultProcessor, interval time.Duration, c chan []Stat) {
	tick := time.Tick(interval * time.Second)
	for {
		select {
		case <-tick:
			d, err := processor.Process()
			if err != nil {
				log.Println(err)
			}
			if d != nil {
				c <- d
			}
		}
	}
}
