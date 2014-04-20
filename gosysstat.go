package main

import (
	"fmt"
	"log"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/processor/cpu"
)

func main() {
	// fetch /proc/diskstats (disk)
	// fetch /proc/meminfo (mem)

	cpuStatProcessor, err := cpu.NewProcessor("/proc/stat")
	if err != nil {
		log.Fatalf(err.Error())
	}
	cpuStatResults := make(chan *[]core.Stat)

	go core.StatProcessor(cpuStatProcessor, cpuStatResults)
	for {
		select {
		case c := <-cpuStatResults:
			fmt.Println(c)
		}
	}
}
