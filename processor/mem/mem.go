package mem

import (
	"github.com/ossareh/gosysstat/core"
	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

type MemStat struct {
	region string
	value  float64
}

func (m *MemStat) Type() string {
	return m.region
}

func (m *MemStat) Values() []float64 {
	return []float64{m.value}
}

type MemProcessor struct {
	rr *reader.ResettingReader
}

func processStatLine(data []string, memTotal, swapTotal float64) *MemStat {
	switch data[0] {
	case "MemTotal:":
		return &MemStat{"total", processor.Atof(data[1])}
	case "MemFree:":
		used := memTotal - processor.Atof(data[1])
		return &MemStat{"used", used}
	case "Cached:":
		return &MemStat{"cached", processor.Atof(data[1])}
	case "SwapTotal:":
		return &MemStat{"swap_total", processor.Atof(data[1])}
	case "SwapFree:":
		used := swapTotal - processor.Atof(data[1])
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
	var memTotal, swapTotal float64
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

func NewProcessor(src reader.DataSource) processor.Processor {
	return &MemProcessor{reader.NewResettingReader(src)}
}
