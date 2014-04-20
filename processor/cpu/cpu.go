package cpu

import (
	"strings"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

type CpuProcessor struct {
	rr *reader.ResettingReader
}

func makeCpuStat(data []string) []int64 {
	return []int64{
		processor.Stoi64(data[0]), // user
		processor.Stoi64(data[1]), // nice
		processor.Stoi64(data[2]), // sys
		processor.Stoi64(data[3]), // idle
		processor.Stoi64(data[4]), // io
	}
}

func processStatLine(data []string) (core.Stat, int) {
	var res core.Stat
	switch data[0] {
	case "cpu":
		res = core.Stat{"total", makeCpuStat(data[1:6])}
	case "intr":
		res = core.Stat{"intr", []int64{processor.Stoi64(data[1])}}
	case "ctxt":
		res = core.Stat{"ctxt", []int64{processor.Stoi64(data[1])}}
	case "processes":
		res = core.Stat{"procs", []int64{processor.Stoi64(data[1])}}
	case "procs_running":
		res = core.Stat{"procsr", []int64{processor.Stoi64(data[1])}}
	case "procs_blocked":
		res = core.Stat{"procsb", []int64{processor.Stoi64(data[1])}}
	default:
		if data[0][0:3] == "cpu" {
			if cpuN := strings.Split(data[0], "cpu")[1]; cpuN != "" {
				res = core.Stat{cpuN, makeCpuStat(data[1:6])}
			}
		} else {
			return res, processor.SKIP
		}
	}
	return res, processor.CONTINUE
}

func (cp CpuProcessor) Process() ([]core.Stat, error) {
	data, err := cp.rr.Read()
	if err != nil {
		// errors may happen...
		return nil, err
	}
	result := []core.Stat{}
	for _, d := range data {
		if len(d) > 0 {
			r, state := processStatLine(d)
			if state != processor.SKIP {
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
