package cpu

import (
	"reflect"
	"testing"

	"github.com/ossareh/libgosysstat/core"
	lt "github.com/ossareh/libgosysstat/processor/testing"
)

func TestCpuProcessor(t *testing.T) {
	th, err := lt.MakeTestHarness("./cpu_sample_one")
	if err != nil {
		t.Fatalf(err.Error())
	}
	proc := NewProcessor(th)
	th.Close()
	if err := th.ReplaceFileHandle("./cpu_sample_two"); err != nil {
		t.Fatalf(err.Error())
	}
	defer th.Close()
	results, _ := proc.Process()
	known := []core.Stat{
		&TotalCpuStat{&CpuStat{91, 2, 206, 7634, 19}},
		&CpuInstanceStat{0, &CpuStat{10, 0, 30, 934, 18}},
		&CpuInstanceStat{1, &CpuStat{16, 1, 28, 948, 1}},
		&CpuInstanceStat{2, &CpuStat{10, 1, 31, 954, 0}},
		&CpuInstanceStat{3, &CpuStat{7, 0, 22, 967, 0}},
		&CpuInstanceStat{4, &CpuStat{1, 0, 6, 991, 0}},
		&CpuInstanceStat{5, &CpuStat{1, 0, 15, 982, 0}},
		&CpuInstanceStat{6, &CpuStat{3, 0, 15, 980, 0}},
		&CpuInstanceStat{7, &CpuStat{44, 0, 58, 879, 0}},
		&SingleStat{"intr", 155844},
		&SingleStat{"ctxt", 1346187},
		&SingleStat{"procs", 3},
		&SingleStat{"procsr", 1},
		&SingleStat{"procsb", 0},
	}

	for i, s := range known {
		if !reflect.DeepEqual(s, results[i]) {
			t.Fatalf("Expected item %d to be %s, got %s", i, s, results[i])
		}
	}
}

func TestCpuStatDelta(t *testing.T) {
	firstSample := []uint64{10, 10, 10, 50, 10}
	secondSample := []uint64{21, 10, 16, 142, 20}
	expected := []uint64{11, 0, 6, 92, 10}
	result := cpuStatDelta(secondSample, firstSample).values()
	if !reflect.DeepEqual(expected, result) {
		for i, e := range expected {
			t.Fatalf("Expected item %d to be %d, got %d", i, e, result[i])
		}
	}
}
