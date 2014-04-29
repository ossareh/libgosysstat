package cpu

import (
	"strconv"
	"strings"

	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
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

func cpuStatDelta(current, initial []float64) *CpuStat {
	user := current[0] - initial[0]
	nice := current[1] - initial[1]
	sys := current[2] - initial[2]
	idle := current[3] - initial[3]
	io := current[4] - initial[4]
	total := user + nice + sys + idle + io

	return &CpuStat{
		user / total,
		nice / total,
		sys / total,
		idle / total,
		io / total,
	}
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
		case "total":
			stat = &TotalCpuStat{cpuStatDelta(res.Values(), prev.Values())}
		case "intr":
			stat = &SingleStat{res.Type(), res.Values()[0] - prev.Values()[0]}
		case "ctxt":
			stat = &SingleStat{res.Type(), res.Values()[0] - prev.Values()[0]}
		case "procs":
			stat = &SingleStat{res.Type(), res.Values()[0] - prev.Values()[0]}
		case "procsr":
			stat = res
		case "procsb":
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

func NewProcessor(src reader.DataSource) processor.Processor {
	cp := &CpuProcessor{reader.NewResettingReader(src), nil}
	cp.Process()
	return cp
}
