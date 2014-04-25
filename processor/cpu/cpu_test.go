package cpu

import (
	"reflect"
	"testing"

	"github.com/ossareh/gosysstat/core"
)

func TestInvalidCpuProcessor(t *testing.T) {
	_, err := NewProcessor("./idontexist")
	if err == nil {
		t.Fatalf("Expected failure opening file", err)
	}
}

func makeCpuMap(user, nice, sys, idle, io int) map[string]int {
	return map[string]int{
		"user": user,
		"nice": nice,
		"sys":  sys,
		"idle": idle,
		"io":   io,
	}
}

func TestCpuProcessor(t *testing.T) {
	proc, err := NewProcessor("./proc_stat.example")
	if err != nil {
		t.Fatalf("Expected to be able to open example file")
	}
	results, _ := proc.Process()
	known := []core.Stat{
		&TotalCpuStat{&CpuStat{488210, 553716, 185158, 155133921, 352874}},
		&CpuInstanceStat{0, &CpuStat{94569, 68276, 55416, 18892780, 317626}},
		&CpuInstanceStat{1, &CpuStat{89363, 70644, 31545, 19393879, 15210}},
		&CpuInstanceStat{2, &CpuStat{88559, 71599, 27731, 19410267, 6418}},
		&CpuInstanceStat{3, &CpuStat{86345, 72636, 26398, 19414920, 4139}},
		&CpuInstanceStat{4, &CpuStat{33012, 65906, 12119, 19503394, 2521}},
		&CpuInstanceStat{5, &CpuStat{33579, 67627, 10803, 19505022, 2238}},
		&CpuInstanceStat{6, &CpuStat{31831, 68442, 10736, 19506646, 1844}},
		&CpuInstanceStat{7, &CpuStat{30947, 68582, 10405, 19507011, 2872}},
		&SingleIntStat{"intr", 122368175},
		&SingleIntStat{"ctxt", 217868872},
		&SingleIntStat{"procs", 6704},
		&SingleIntStat{"procsr", 1},
		&SingleIntStat{"procsb", 0},
	}

	if !reflect.DeepEqual(known, results) {
		t.Fatalf("Expected matching results", known, results)
	}
}
