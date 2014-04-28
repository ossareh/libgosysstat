package cpu

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ossareh/gosysstat/core"
	lt "github.com/ossareh/gosysstat/processor/testing"
)

func TestCpuProcessor(t *testing.T) {
	th, err := lt.MakeTestHarness("./proc_stat.example")
	if err != nil {
		t.Fatalf(err.Error())
	}
	proc := NewProcessor(th)
	defer th.Close()
	/*if err := th.replaceFileHandle("./cpu_sample_two"); err != nil {
		t.Fatalf(err.Error())
	}
	defer th.Close()*/
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
		&SingleStat{"intr", 122368175},
		&SingleStat{"ctxt", 217868872},
		&SingleStat{"procs", 6704},
		&SingleStat{"procsr", 1},
		&SingleStat{"procsb", 0},
	}

	if !reflect.DeepEqual(known, results) {
		t.Fatalf("Expected matching results", known, results)
	}
}

func TestCpuStatSubtract(t *testing.T) {
	firstSample := &CpuStat{10, 10, 10, 50, 10}
	secondSample := &CpuStat{21, 10, 16, 142, 20}
	result := secondSample.Subtract(firstSample)
	e := "0.09"
	if r := fmt.Sprintf("%.2f", result.user); r != e {
		t.Fatalf("Expected %d got %s", e, r)
	}
	e = "0.00"
	if r := fmt.Sprintf("%.2f", result.nice); r != e {
		t.Fatalf("Expected %d got %s", e, r)
	}
	e = "0.05"
	if r := fmt.Sprintf("%.2f", result.sys); r != e {
		t.Fatalf("Expected %d got %s", e, r)
	}
	e = "0.77"
	if r := fmt.Sprintf("%.2f", result.idle); r != e {
		t.Fatalf("Expected %d got %s", e, r)
	}
	e = "0.08"
	if r := fmt.Sprintf("%.2f", result.io); r != e {
		t.Fatalf("Expected %d got %s", e, r)
	}
}
