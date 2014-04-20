package core

import "time"

type Stat struct {
	Type   string
	Values []int64
}

type ResultProcessor interface {
	Process() ([]Stat, error)
}

func StatProcessor(processor ResultProcessor, c chan *[]Stat) {
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-tick:
			d, err := processor.Process()
			if err == nil {
				c <- &d
			} else {
				// TODO: Dont Panic!
				panic(err)
			}
		}
	}
}
