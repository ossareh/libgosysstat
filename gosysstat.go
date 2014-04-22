package main

import (
	"fmt"
	"log"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/processor/cpu"
	"github.com/ossareh/gosysstat/processor/mem"
)

const TICK_INTERVAL = 1

func main() {
	// fetch /proc/diskstats (disk)
	// fetch /proc/meminfo (mem)

	cpuStatProcessor, err := cpu.NewProcessor("/proc/stat")
	if err != nil {
		log.Fatalf(err.Error())
	}
	cpuStatResults := make(chan []core.Stat)

	memStatProcessor, err := mem.NewProcessor("/proc/meminfo")
	if err != nil {
		log.Fatalf(err.Error())
	}
	memStatResults := make(chan []core.Stat)

	go core.StatProcessor(cpuStatProcessor, TICK_INTERVAL, cpuStatResults)
	go core.StatProcessor(memStatProcessor, TICK_INTERVAL, memStatResults)
	for {
		select {
		case c := <-cpuStatResults:
			for _, s := range c {
				fmt.Println(s.Type())
				fmt.Println(s.Values())
			}
		case c := <-memStatResults:
			fmt.Println(c)
		}
	}
}
