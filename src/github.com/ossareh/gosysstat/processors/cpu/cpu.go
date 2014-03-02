package cpu

import (
	"github.com/ossareh/gosysstat/core"
)

const (
	StatFile = "/proc/stat"
)

/*
  FSM:
  if STATE == READING_CONTEXT && bytes_read < 3:
    READ_BYTE

  if STATE == READING_PROCESS_CONTEXT && bytes_read < 7:
    READ_BYTE

  if STATE == READING_CPU_CONTEXT && bytes_read < '4':
    READ_BYTE

  if STATE == READING_CONTEXT && bytes_read == 3:
    if bytes_read = ['c', 'p', 'u']:
      STATE = READING_CPU_CONTEXT
    if bytes_read = ['i', 'n', 't']:
      STATE = READING_INTERUPTS
    if bytes_read = ['c', 't', 'x']
      STATE = READING_CONTEXT_SWITCHES
    if bytes_read = ['p', 'r', 'o']
      STATE = READING_PROCESS_CONTEXT

  if STATE == READING_PROCESS_CONTEXT && bytes_read == 7:
    switch byte[6]:
    case 'r':
      STATE = READING_PROCESSES_RUNNING
    case 'b':
      STATE = READING_PROCESSES_BLOCKED
    default:
      STATE = READING_PROCESSES_TOTAL
*/

func ProcessBytes(d []byte) []core.Stat {
	stats := make([]core.Stat, 7)


	for i := 0; i < 3; i++ {
		stats[i] = core.Stat{"CPU", "TOTAL", int64(i*i)}
	}
	return stats
}
