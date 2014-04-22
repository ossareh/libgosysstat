package cpu

import (
	"strconv"

	"github.com/ossareh/gosysstat/core/reader"
	"github.com/ossareh/gosysstat/processor"
)

type CpuStat struct {
	user int
	nice int
	sys  int
	idle int
	io   int
}

type TotalCpuStat struct {
	stats *CpuStat
}

func (t *TotalCpuStat) Type() string {
	return "total"
}

func (t *TotalCpuStat) Values() []int {
	return []int{
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

func (t *CpuInstanceStat) Values() []int {
	return []int{
		t.stats.user,
		t.stats.nice,
		t.stats.sys,
		t.stats.idle,
		t.stats.io,
	}
}

type SingleIntStat struct {
	kind  string
	value int
}

func (s *SingleIntStat) Type() string {
	return s.kind
}

func (s *SingleIntStat) Values() []int {
	return []int{s.value}
}

type CpuProcessor struct {
	rr *reader.ResettingReader
}

func makeCpuStat(data []string) *CpuStat {
	return &CpuStat{
		processor.Stoi(data[0]), // user
		processor.Stoi(data[1]), // nice
		processor.Stoi(data[2]), // sys
		processor.Stoi(data[3]), // idle
		processor.Stoi(data[4]), // io
	}
}
