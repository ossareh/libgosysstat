package mem

import (
	"reflect"
	"testing"

	"github.com/ossareh/gosysstat/core"
	lt "github.com/ossareh/gosysstat/processor/testing"
)

func TestMemProcessor(t *testing.T) {
	th, err := lt.MakeTestHarness("./proc_meminfo.example")
	if err != nil {
		t.Fatalf(err.Error())
	}
	proc := NewProcessor(th)
	defer th.Close()
	/*if err := th.ReplaceFileHandle("./proc_meminfo.example2"); err != nil {
		t.Fatalf(err.Error())
	}
	defer th.Close()*/
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
