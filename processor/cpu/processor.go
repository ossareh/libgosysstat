package cpu

import (
	"strconv"
	"strings"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

func processStatLine(data []string) (core.Stat, error) {
	switch data[0] {
	case "cpu":
		return &TotalCpuStat{makeCpuStat(data[1:6])}, nil
	case "intr":
		return &SingleIntStat{"intr", processor.Stoi(data[1])}, nil
	case "ctxt":
		return &SingleIntStat{"ctxt", processor.Stoi(data[1])}, nil
	case "processes":
		return &SingleIntStat{"procs", processor.Stoi(data[1])}, nil
	case "procs_running":
		return &SingleIntStat{"procsr", processor.Stoi(data[1])}, nil
	case "procs_blocked":
		return &SingleIntStat{"procsb", processor.Stoi(data[1])}, nil
	default:
		if data[0][0:3] == "cpu" {
			if cpuN := strings.Split(data[0], "cpu")[1]; cpuN != "" {
				cpuN, err := strconv.Atoi(cpuN)
				if err != nil {
					return nil, err
				}
				return &CpuInstanceStat{cpuN, makeCpuStat(data[1:6])}, nil
			}
		}
		return nil, nil
	}
}

func (cp CpuProcessor) Process() ([]core.Stat, error) {
	data, err := cp.rr.Read()
	if err != nil {
		return nil, err
	}
	result := []core.Stat{}
	for _, d := range data {
		if len(d) > 0 {
			r, err := processStatLine(d)
			if err != nil {
				return nil, err
			}
			if r != nil {
				result = append(result, r)
			}
		}
	}
	return result, nil
}

func NewProcessor(filename string) (processor.Processor, error) {
	rr, err := reader.Open(filename)
	if err != nil {
		return nil, err
	}
	return &CpuProcessor{rr}, nil
}
