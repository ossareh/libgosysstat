package cpu

import (
	"github.com/ossareh/gosysstat/core"
)

const (
	StatFile = "/proc/stat"
)

func ProcessBytes(d []byte) []core.Stat {
	stats := make([]core.Stat, 3)
	for i := 0; i < 3; i++ {
		stats[i] = core.Stat{"CPU", "TOTAL", int64(i*i)}
	}
	return stats
}
