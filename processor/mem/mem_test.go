package mem

import (
	"reflect"
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
		&MemStat{"total", 24684748},
		&MemStat{"used", 4749476},
		&MemStat{"cached", 1919332},
		&MemStat{"swap_total", 0},
		&MemStat{"swap_free", 0},
	}

	if !reflect.DeepEqual(known, results) {
		t.Fatalf("Expected matching results", known, results)
	}
}
