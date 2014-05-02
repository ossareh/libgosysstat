package processor

import (
	"log"
	"strconv"

	"github.com/ossareh/libgosysstat/core"
)

type Processor interface {
	Process() ([]core.Stat, error)
}

func Atoui64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return i
}
