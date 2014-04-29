package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/processor/cpu"
)

const (
	TICK_INTERVAL       = 1
	CPU_STAT_FMT        = "%s:(user:%.2f%%, sys:%.2f%%, idle:%.2f%%, io:%.2f%%)\n"
	CPU_SINGLE_STAT_FMT = "%s:%d\n"
)

func prepareCpuValues(values []float64) (user, sys, idle, io float64) {
	user = (values[0] + values[1]) * 100
	sys = values[2] * 100
	idle = values[3] * 100
	io = values[4] * 100
	return
}

func formatCpuStat(data []core.Stat) string {
	var buf bytes.Buffer

	for _, d := range data {
		values := d.Values()
		var s string
		switch d.Type() {
		case "total":
			user, sys, idle, io := prepareCpuValues(values)
			s = fmt.Sprintf(CPU_STAT_FMT, "Total", user, sys, idle, io)
		case "intr":
			s = fmt.Sprintf(CPU_SINGLE_STAT_FMT, "Interupts", int(values[0]))
		case "ctxt":
			s = fmt.Sprintf(CPU_SINGLE_STAT_FMT, "Context Switches", int(values[0]))
		case "procs":
		case "procsr":
		case "procsb":
		default:
			user, sys, idle, io := prepareCpuValues(values)
			s = fmt.Sprintf(CPU_STAT_FMT, "CPU"+d.Type(), user, sys, idle, io)
		}
		buf.WriteString(s)
	}
	return buf.String()
}

func main() {
	// fetch /proc/diskstats (disk)
	// fetch /proc/meminfo (mem)

	cpuFh, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer cpuFh.Close()
	cpuStatProcessor := cpu.NewProcessor(cpuFh)
	cpuStatResults := make(chan []core.Stat)

	/*	memFh, err := os.Open("/proc/stat")
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer memFh.Close()
		memStatProcessor := mem.NewProcessor(memFh)
		memStatResults := make(chan []core.Stat)*/

	go core.StatProcessor(cpuStatProcessor, TICK_INTERVAL, cpuStatResults)
	//go core.StatProcessor(memStatProcessor, TICK_INTERVAL, memStatResults)
	for {
		select {
		case c := <-cpuStatResults:
			fmt.Println(formatCpuStat(c))

			/*case c := <-memStatResults:
			fmt.Println(c)*/
		}
	}
}
