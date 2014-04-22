package cpu

import (
	"fmt"
	"testing"

	"github.com/ossareh/gosysstat/core"
)

func TestInvalidCpuProcessor(t *testing.T) {
	_, err := NewProcessor("./idontexist")
	if err == nil {
		t.Fatalf("Expected failure opening file", err)
	}
}

func makeCpuMap(user, nice, sys, idle, io int64) map[string]int64 {
	return map[string]int64{
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
		core.Stat{"total", []int64{488210, 553716, 185158, 155133921, 352874}},
		core.Stat{"0", []int64{94569, 68276, 55416, 18892780, 317626}},
		core.Stat{"1", []int64{89363, 70644, 31545, 19393879, 15210}},
		core.Stat{"2", []int64{88559, 71599, 27731, 19410267, 6418}},
		core.Stat{"3", []int64{86345, 72636, 26398, 19414920, 4139}},
		core.Stat{"4", []int64{33012, 65906, 12119, 19503394, 2521}},
		core.Stat{"5", []int64{33579, 67627, 10803, 19505022, 2238}},
		core.Stat{"6", []int64{31831, 68442, 10736, 19506646, 1844}},
		core.Stat{"7", []int64{30947, 68582, 10405, 19507011, 2872}},
		core.Stat{"intr", []int64{122368175}},
		core.Stat{"ctxt", []int64{217868872}},
		core.Stat{"procs", []int64{6704}},
		core.Stat{"procsr", []int64{1}},
		core.Stat{"procsb", []int64{0}},
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
