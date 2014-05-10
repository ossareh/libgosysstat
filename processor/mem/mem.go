package mem

import (
	"github.com/ossareh/libgosysstat/core"
	"github.com/ossareh/libgosysstat/core/reader"
	"github.com/ossareh/libgosysstat/processor"
)

type MemStat struct {
	region string
	value  uint64
}

func (m *MemStat) Type() string {
	return m.region
}

func (m *MemStat) Values() []uint64 {
	return []uint64{m.value}
}

type MemProcessor struct {
	rr *reader.ResettingReader
}

func processStatLine(data []string, memTotal, swapTotal uint64) *MemStat {
	switch data[0] {
	case "MemTotal:":
		return &MemStat{"total", processor.Atoui64(data[1])}
	case "MemFree:":
		used := memTotal - processor.Atoui64(data[1])
		return &MemStat{"used", used}
	case "Cached:":
		return &MemStat{"cached", processor.Atoui64(data[1])}
	case "SwapTotal:":
		return &MemStat{"swap_total", processor.Atoui64(data[1])}
	case "SwapFree:":
		used := swapTotal - processor.Atoui64(data[1])
		return &MemStat{"swap_used", used}
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
	var memTotal, swapTotal uint64
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

func New(src reader.DataSource) processor.Processor {
	return &MemProcessor{reader.New(src)}
}
