package cpu

import (
	"strconv"

	"github.com/ossareh/gosysstat/processor"
)

type CpuStat struct {
	user float64
	nice float64
	sys  float64
	idle float64
	io   float64
}

func (cs *CpuStat) Subtract(previous *CpuStat) *CpuStat {
	user := cs.user - previous.user
	nice := cs.nice - previous.nice
	sys := cs.sys - previous.sys
	idle := cs.idle - previous.idle
	io := cs.io - previous.io
	total := user + nice + sys + idle + io
	return &CpuStat{
		user / total,
		nice / total,
		sys / total,
		idle / total,
		io / total,
	}
}

type TotalCpuStat struct {
	stats *CpuStat
}

func (t *TotalCpuStat) Type() string {
	return "total"
}

func (t *TotalCpuStat) Values() []float64 {
	return []float64{
		t.stats.user,
		t.stats.nice,
		t.stats.sys,
		t.stats.idle,
		t.stats.io,
	}
}

type CpuInstanceStat struct {
	cpuInstance int
	stats       *CpuStat
}

func (t *CpuInstanceStat) Type() string {
	return strconv.Itoa(t.cpuInstance)
}

func (t *CpuInstanceStat) Values() []float64 {
	return []float64{
		t.stats.user,
		t.stats.nice,
		t.stats.sys,
		t.stats.idle,
		t.stats.io,
	}
}

type SingleStat struct {
	kind  string
	value float64
}

func (s *SingleStat) Type() string {
	return s.kind
}

func (s *SingleStat) Values() []float64 {
	return []float64{s.value}
}

func makeCpuStat(data []string) *CpuStat {
	return &CpuStat{
		processor.Atof(data[0]), // user
		processor.Atof(data[1]), // nice
		processor.Atof(data[2]), // sys
		processor.Atof(data[3]), // idle
		processor.Atof(data[4]), // io
	}
}
