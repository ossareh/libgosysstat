package processor

import (
	"strconv"

	"github.com/ossareh/gosysstat/core"
)

const (
	CONTINUE = iota
	SKIP
	STOP
)

type Processor interface {
	Process() ([]core.Stat, error)
}

func Stoi64(s string) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i64
}
