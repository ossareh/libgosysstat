package processor

import (
	"log"
	"strconv"

	"github.com/ossareh/libgosysstat/core"
)

type Processor interface {
	Process() ([]core.Stat, error)
}

func Atof(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return i
}
