package mem

import (
	"fmt"
	"testing"

	"github.com/ossareh/gosysstat/core"
)

func TestInvalidMemProcessor(t *testing.T) {
	_, err := NewProcessor("./idontexist")
	if err == nil {
		t.Fatalf("Expected failure opening file", err)
	}
}

func TestMemProcessor(t *testing.T) {
	proc, err := NewProcessor("./proc_meminfo.example")
	if err != nil {
		t.Fatalf("Expected to be able to open example file")
	}
	results, _ := proc.Process()
	known := []core.Stat{
		core.Stat{"total", []int64{24684748}},
		core.Stat{"used", []int64{4749476}},
		core.Stat{"cached", []int64{}},
		core.Stat{"swap_total", []int64{}},
		core.Stat{"swap_free", []int64{}},
	}

	if len(known) != len(results) {
		t.Fatalf("Expected same number of results", len(known), len(results))
	}

	for idx, stat := range known {
		if stat.Type != results[idx].Type {
			t.Fatalf("Stat Type mismatch", stat.Type, results[idx].Type)
		}
		for subIdx, val := range stat.Values {
			if val != results[idx].Values[subIdx] {
				s := fmt.Sprintf("Stat Value mismatch (item:%d, val:%d)",
					idx, subIdx)
				t.Fatalf(s, val, results[idx].Values[subIdx])
			}
		}
	}
}
