package cpu

import (
	"testing"
	"io/ioutil"
	"github.com/ossareh/gosysstat/core"
)

func TestStatFile(t *testing.T) {
	if StatFile != "/proc/stat" {
		t.Fail()
	}
}

func makeTestStats() []core.Stat {
	stats := make([]core.Stat, 14)
	stats[0]  = core.Stat{"CPU", "Total",   488210}
	stats[1]  = core.Stat{"CPU", "0",        94569}
	stats[2]  = core.Stat{"CPU", "1",        89363}
	stats[3]  = core.Stat{"CPU", "2",        88559}
	stats[4]  = core.Stat{"CPU", "3",        86345}
	stats[5]  = core.Stat{"CPU", "4",        33012}
	stats[6]  = core.Stat{"CPU", "5",        33579}
	stats[7]  = core.Stat{"CPU", "6",        31831}
	stats[8]  = core.Stat{"CPU", "7",        30947}
	stats[9]  = core.Stat{"CPU", "intr", 122368175}
	stats[10] = core.Stat{"CPU", "ctxt", 217868872}
	stats[11] = core.Stat{"CPU", "procs",     6704}
	stats[12] = core.Stat{"CPU", "procsr",       1}
	stats[13] = core.Stat{"CPU", "procsb",       0}
	return stats
}

func TestProcessBytes(t *testing.T) {
	fixture_bytes, err := ioutil.ReadFile("./cpu_stat.example")
	if err != nil { t.Error(err) }
	stats := ProcessBytes(fixture_bytes)
	expectedStats := makeTestStats()
	if len(stats) != len(expectedStats) {
		t.Fail()
	}
	for idx, stat := range expectedStats {
		if stats[idx].Type != stat.Type {
			t.Fail()
		}
		if stats[idx].Aspect != stat.Aspect {
			t.Fail()
		}
		if stats[idx].Value != stat.Value {
			t.Fail()
		}
	}
}
