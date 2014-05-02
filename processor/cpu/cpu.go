package cpu

import (
	"strconv"

	"github.com/ossareh/libgosysstat/processor"
)

const (
	TOTAL         string = "total"
	INTR                 = "intr"
	CTXT                 = "ctxt"
	PROCS                = "procs"
	PROCS_RUNNING        = "procsr"
	PROCS_BLOCKED        = "procsb"
)

type CpuStat struct {
	user uint64
	nice uint64
	sys  uint64
	idle uint64
	io   uint64
}

func (c *CpuStat) values() []uint64 {
	return []uint64{
		c.user,
		c.nice,
		c.sys,
		c.idle,
		c.io,
	}
}

type TotalCpuStat struct {
	stats *CpuStat
}

func (t *TotalCpuStat) Type() string {
	return TOTAL
}

func (t *TotalCpuStat) Values() []uint64 {
	return t.stats.values()
}

type CpuInstanceStat struct {
	cpuInstance int
	stats       *CpuStat
}

func (t *CpuInstanceStat) Type() string {
	return strconv.Itoa(t.cpuInstance)
}

func (t *CpuInstanceStat) Values() []uint64 {
	return t.stats.values()
}

type SingleStat struct {
	kind  string
	value uint64
}

func (s *SingleStat) Type() string {
	return s.kind
}

func (s *SingleStat) Values() []uint64 {
	return []uint64{s.value}
}

func makeCpuStat(data []string) *CpuStat {
	return &CpuStat{
		processor.Atoui64(data[0]), // user
		processor.Atoui64(data[1]), // nice
		processor.Atoui64(data[2]), // sys
		processor.Atoui64(data[3]), // idle
		processor.Atoui64(data[4]), // io
	}
}
