package cpu

import (
	"strconv"
	"strings"
	"time"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

type ProcessResult struct {
	result []core.Stat
	when   time.Time
}

type CpuProcessor struct {
	rr             *reader.ResettingReader
	previousResult *ProcessResult
}

func processStatLine(data []string) (core.Stat, error) {
	switch data[0] {
	case "cpu":
		return &TotalCpuStat{makeCpuStat(data[1:6])}, nil
	case "intr":
		return &SingleStat{"intr", processor.Atof(data[1])}, nil
	case "ctxt":
		return &SingleStat{"ctxt", processor.Atof(data[1])}, nil
	case "processes":
		return &SingleStat{"procs", processor.Atof(data[1])}, nil
	case "procs_running":
		return &SingleStat{"procsr", processor.Atof(data[1])}, nil
	case "procs_blocked":
		return &SingleStat{"procsb", processor.Atof(data[1])}, nil
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

func readStats(cp *CpuProcessor) ([]core.Stat, error) {
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

func (cp *CpuProcessor) Process() ([]core.Stat, error) {
	rawStats, err := readStats(cp)
	if err != nil {
		return nil, err
	}
	result := &ProcessResult{rawStats, time.Now()}
	if cp.previousResult != nil {

	}
	cp.previousResult = result
	return rawStats, nil
}

func NewProcessor(src reader.DataSource) processor.Processor {
	cp := &CpuProcessor{reader.NewResettingReader(src), nil}
	cp.Process()
	return cp
}
