package core

import (
	"io/ioutil"
	"time"
)

type Stat struct {
	Type   string
	Aspect string
	Values map[string]int64
}

type FileProcessor func(string) []Stat

func readFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func StatProcessor(filename string, processor FileProcessor, c chan *[]Stat) {
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-tick:
			data := string(readFile(filename))
			stats := processor(data)
			c <- &stats
		}
	}
}
