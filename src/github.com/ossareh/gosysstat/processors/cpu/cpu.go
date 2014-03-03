package cpu

import (
	"strconv"
	"strings"

	"github.com/ossareh/gosysstat/core"
)

const (
	StatFile = "/proc/stat"
)

func convertToInts(input []string) []int64 {
	output := make([]int64, len(input))
	for idx, s := range input {
		i, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			panic(err)
		}
		output[idx] = i
	}
	return output
}

func convertToInt(input string) int64 {
	return convertToInts([]string{input})[0]
}

func makeCpuValuesMap(input []int64) map[string]int64 {
	return map[string]int64{
		"user": input[0],
		"nice": input[1],
		"sys":  input[2],
		"idle": input[3],
		"io":   input[4]}
}

func ProcessBytes(data string) []core.Stat {
	// TODO: improve
	stats := make([]core.Stat, 0)
	context := "CPU"
	processesValues := make(map[string]int64, 3)
	// TODO: this can probably be a stack of start/end positions for
	// interesting data and then we can use slices against a byte array
	// rather than creating the string that we'll just throw away
	for _, line := range strings.Split(data, "\n") {
		fields := strings.Fields(line)
		var valueMap map[string]int64
		var aspect string

		if len(fields) == 0 {
			break
		}
		switch fields[0] {
		case "cpu":
			aspect = "total"
			valueMap = makeCpuValuesMap(convertToInts(fields[1:6]))
		case "intr":
			aspect = "intr"
			valueMap = map[string]int64{
				"total": convertToInt(fields[1])}
		case "ctxt":
			aspect = "ctxt"
			valueMap = map[string]int64{
				"total": convertToInt(fields[1])}
		case "processes":
			processesValues["total"] = convertToInt(fields[1])
		case "procs_running":
			processesValues["running"] = convertToInt(fields[1])
		case "procs_blocked":
			processesValues["blocked"] = convertToInt(fields[1])
		default:
			if fields[0][0:3] == "cpu" {
				if cpuN := strings.Split(fields[0], "cpu")[1]; cpuN != "" {
					aspect = cpuN
					valueMap = makeCpuValuesMap(convertToInts(fields[1:6]))
				}
			}
		}

		if valueMap != nil {
			stats = append(stats, core.Stat{context, aspect, valueMap})
		}
	}
	stats = append(stats, core.Stat{context, "procs", processesValues})
	return stats
}
