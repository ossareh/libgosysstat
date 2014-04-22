package processor

import (
	"log"
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

func Stoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return i
}
