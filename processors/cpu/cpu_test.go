package cpu

import (
	"io/ioutil"
	"testing"

	"github.com/ossareh/gosysstat/core"
)

func TestStatFile(t *testing.T) {
	if StatFile != "/proc/stat" {
		t.Fail()
	}
}

func makeCpuMap(user, nice, sys, idle, io int64) map[string]int64 {
	return map[string]int64{
		"user": user,
		"nice": nice,
		"sys":  sys,
		"idle": idle,
		"io":   io}
}

func makeTestStats() []core.Stat {
	stats := make([]core.Stat, 12)
	stats[0] = core.Stat{"CPU", "total",
		makeCpuMap(488210, 553716, 185158, 155133921, 352874)}

	stats[1] = core.Stat{"CPU", "0",
		makeCpuMap(94569, 68276, 55416, 18892780, 317626)}

	stats[2] = core.Stat{"CPU", "1",
		makeCpuMap(89363, 70644, 31545, 19393879, 15210)}

	stats[3] = core.Stat{"CPU", "2",
		makeCpuMap(88559, 71599, 27731, 19410267, 6418)}

	stats[4] = core.Stat{"CPU", "3",
		makeCpuMap(86345, 72636, 26398, 19414920, 4139)}

	stats[5] = core.Stat{"CPU", "4",
		makeCpuMap(33012, 65906, 12119, 19503394, 2521)}

	stats[6] = core.Stat{"CPU", "5",
		makeCpuMap(33579, 67627, 10803, 19505022, 2238)}

	stats[7] = core.Stat{"CPU", "6",
		makeCpuMap(31831, 68442, 10736, 19506646, 1844)}

	stats[8] = core.Stat{"CPU", "7",
		makeCpuMap(30947, 68582, 10405, 19507011, 2872)}

	stats[9] = core.Stat{"CPU", "intr", map[string]int64{
		"total": 122368175}}

	stats[10] = core.Stat{"CPU", "ctxt", map[string]int64{
		"total": 217868872}}

	stats[11] = core.Stat{"CPU", "procs", map[string]int64{
		"total":   6704,
		"running": 1,
		"blocked": 0}}

	return stats
}

func TestProcessData(t *testing.T) {
	fixture_bytes, err := ioutil.ReadFile("./cpu_stat.example")
	if err != nil {
		t.Error(err)
	}
	stats := ProcessData(string(fixture_bytes))
	expectedStats := makeTestStats()
	if len(stats) != len(expectedStats) {
		t.Error("%v != %v", stats, expectedStats)
	}
	for idx, stat := range expectedStats {
		if stats[idx].Type != stat.Type {
			t.Errorf("%v != %v", stats[idx].Type, stat.Type)
		}
		if stats[idx].Aspect != stat.Aspect {
			t.Errorf("%v != %v", stats[idx].Aspect, stat.Aspect)
		}
		for k := range stats[idx].Values {
			if stats[idx].Values[k] != stat.Values[k] {
				t.Errorf("%v != %v", stats[idx].Values[k], stat.Values[k])
			}
		}
	}
}
