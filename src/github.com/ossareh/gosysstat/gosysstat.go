package main

import (
	"fmt"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/processors/cpu"
)

func main() {
	// fetch /proc/diskstats (disk)
	// fetch /proc/meminfo (mem)

	// fetch /proc/stat (cpu)
	cpuStats := make(chan *[]core.Stat)
	go core.StatProcessor(cpu.StatFile, cpu.ProcessBytes, cpuStats)
	for {
		select {
		case c := <-cpuStats:
			fmt.Println(c)
		}
	}
}
