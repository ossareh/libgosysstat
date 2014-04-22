package mem

import (
	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

type MemProcessor struct {
	rr *reader.ResettingReader
}

func processStatLine(data []string, memTotal, swapTotal int) (core.Stat, int) {
	var res core.Stat
	switch data[0] {
	case "MemTotal:":
		res = core.Stat{"total", []int{processor.Stoi(data[1])}}
	case "MemFree:":
		used := memTotal - processor.Stoi(data[1])
		res = core.Stat{"used", []int{used}}
	case "Cached:":
		res = core.Stat{"cached", []int{processor.Stoi(data[1])}}
	case "SwapTotal:":
		res = core.Stat{"swap_total", []int{processor.Stoi(data[1])}}
	case "SwapFree:":
		used := swapTotal - processor.Stoi(data[1])
		res = core.Stat{"swap_free", []int{used}}
	default:
		return res, processor.SKIP
	}
	return res, processor.CONTINUE

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
			r, state := processStatLine(d, memTotal, swapTotal)
			if state != processor.SKIP {
				switch r.Type {
				case "total":
					memTotal = r.Values[0]
				case "swap_total":
					swapTotal = r.Values[0]
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
