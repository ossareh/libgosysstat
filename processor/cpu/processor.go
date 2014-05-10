package cpu

import (
	"strconv"
	"strings"

	"github.com/ossareh/libgosysstat/core"
	"github.com/ossareh/libgosysstat/core/reader"
	"github.com/ossareh/libgosysstat/processor"
)

type CpuProcessor struct {
	rr             *reader.ResettingReader
	previousResult []core.Stat
}

func processStatLine(data []string) (core.Stat, error) {
	switch data[0] {
	case "cpu":
		return &TotalCpuStat{makeCpuStat(data[1:6])}, nil
	case "intr":
		return &SingleStat{INTR, processor.Atoui64(data[1])}, nil
	case "ctxt":
		return &SingleStat{CTXT, processor.Atoui64(data[1])}, nil
	case "processes":
		return &SingleStat{PROCS, processor.Atoui64(data[1])}, nil
	case "procs_running":
		return &SingleStat{PROCS_RUNNING, processor.Atoui64(data[1])}, nil
	case "procs_blocked":
		return &SingleStat{PROCS_BLOCKED, processor.Atoui64(data[1])}, nil
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

func cpuStatDelta(current, initial []uint64) *CpuStat {
	user := current[0] - initial[0]
	nice := current[1] - initial[1]
	sys := current[2] - initial[2]
	idle := current[3] - initial[3]
	io := current[4] - initial[4]

	return &CpuStat{user, nice, sys, idle, io}
}

func (cp *CpuProcessor) Process() ([]core.Stat, error) {
	rawResult, err := readStats(cp)
	if err != nil {
		return nil, err
	}
	if cp.previousResult == nil {
		cp.previousResult = rawResult
		return nil, nil
	}
	computedResult := make([]core.Stat, len(rawResult))
	for idx, res := range rawResult {
		var stat core.Stat
		prev := cp.previousResult[idx]
		switch res.Type() {
		case TOTAL:
			stat = &TotalCpuStat{cpuStatDelta(res.Values(), prev.Values())}
		case INTR:
			stat = &SingleStat{res.Type(), res.Values()[0] - prev.Values()[0]}
		case CTXT:
			stat = &SingleStat{res.Type(), res.Values()[0] - prev.Values()[0]}
		case PROCS:
			stat = &SingleStat{res.Type(), res.Values()[0] - prev.Values()[0]}
		case PROCS_RUNNING:
			stat = res
		case PROCS_BLOCKED:
			stat = res
		default:
			// handling individual cores
			i, _ := strconv.Atoi(res.Type())
			stat = &CpuInstanceStat{i, cpuStatDelta(res.Values(), prev.Values())}
		}
		computedResult[idx] = stat
	}
	cp.previousResult = rawResult
	return computedResult, nil
}

func New(src reader.DataSource) processor.Processor {
	cp := &CpuProcessor{reader.New(src), nil}
	cp.Process()
	return cp
}
