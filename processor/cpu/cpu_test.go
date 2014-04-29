package cpu

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ossareh/gosysstat/core"
	lt "github.com/ossareh/gosysstat/processor/testing"
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
		&TotalCpuStat{&CpuStat{0.011443661971830986, 0.00025150905432595576, 0.02590543259557344, 0.9600100603621731, 0.0023893360160965795}},
		&CpuInstanceStat{0, &CpuStat{0.010080645161290322, 0, 0.03024193548387097, 0.9415322580645161, 0.018145161290322582}},
		&CpuInstanceStat{1, &CpuStat{0.01609657947686117, 0.001006036217303823, 0.028169014084507043, 0.9537223340040242, 0.001006036217303823}},
		&CpuInstanceStat{2, &CpuStat{0.010040160642570281, 0.001004016064257028, 0.03112449799196787, 0.9578313253012049, 0}},
		&CpuInstanceStat{3, &CpuStat{0.007028112449799197, 0, 0.02208835341365462, 0.9708835341365462, 0}},
		&CpuInstanceStat{4, &CpuStat{0.001002004008016032, 0, 0.006012024048096192, 0.9929859719438878, 0}},
		&CpuInstanceStat{5, &CpuStat{0.001002004008016032, 0, 0.01503006012024048, 0.9839679358717435, 0}},
		&CpuInstanceStat{6, &CpuStat{0.003006012024048096, 0, 0.01503006012024048, 0.9819639278557114, 0}},
		&CpuInstanceStat{7, &CpuStat{0.04485219164118247, 0, 0.059123343527013254, 0.8960244648318043, 0}},
		&SingleStat{"intr", 155844},
		&SingleStat{"ctxt", 1346187},
		&SingleStat{"procs", 84711},
		&SingleStat{"procsr", 1},
		&SingleStat{"procsb", 0},
	}

	for i, s := range known {
		if !reflect.DeepEqual(s, results[i]) {
			t.Fatalf("Expected %s, got %s", s, results[i])
		}
	}
}

func TestCpuStatDelta(t *testing.T) {
	firstSample := []float64{10, 10, 10, 50, 10}
	secondSample := []float64{21, 10, 16, 142, 20}
	result := cpuStatDelta(secondSample, firstSample)
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
