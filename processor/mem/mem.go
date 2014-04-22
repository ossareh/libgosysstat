package mem

import (
	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

type MemStat struct {
	region string
	value  int
}

func (m *MemStat) Type() string {
	return m.region
}

func (m *MemStat) Values() []int {
	return []int{m.value}
}

type MemProcessor struct {
	rr *reader.ResettingReader
}

func processStatLine(data []string, memTotal, swapTotal int) *MemStat {
	switch data[0] {
	case "MemTotal:":
		return &MemStat{"total", processor.Stoi(data[1])}
	case "MemFree:":
		used := memTotal - processor.Stoi(data[1])
		return &MemStat{"used", used}
	case "Cached:":
		return &MemStat{"cached", processor.Stoi(data[1])}
	case "SwapTotal:":
		return &MemStat{"swap_total", processor.Stoi(data[1])}
	case "SwapFree:":
		used := swapTotal - processor.Stoi(data[1])
		return &MemStat{"swap_free", used}
	default:
		return nil
	}
}

func (mp *MemProcessor) Process() ([]core.Stat, error) {
	data, err := mp.rr.Read()
	if err != nil {
		return nil, err
	}
	result := []core.Stat{}
	var memTotal, swapTotal int
	for _, d := range data {
		if len(d) > 0 {
			r := processStatLine(d, memTotal, swapTotal)
			if r != nil {
				switch r.region {
				case "total":
					memTotal = r.value
				case "swap_total":
					swapTotal = r.value
				}
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
	return &MemProcessor{rr}, nil
}
