package reader

import (
	"io"
	"os"
)

const (
	INITIAL_ROWS  = 1
	INITIAL_SLICE = 512
)

type ResettingReader struct {
	fh *os.File
}

func tokenize(bytes []byte) [][]string {
	var row, start int
	rows := make([][]string, INITIAL_ROWS)
	rows[row] = []string{}
	for idx, b := range bytes {
		// skip forwards if we're just finding whitespace
		if idx == start && b == ' ' {
			start = idx + 1
			continue
		}

		// handle case where last line does not end with a new line
		if idx == len(bytes)-1 {
			idx = len(bytes)
		}

		// normal case, hit a space or new line, add to rows
		if b == ' ' || b == '\n' || idx == len(bytes) {
			rows[row] = append(rows[row], string(bytes[start:idx]))
			start = idx + 1
		}

		// If we're looking at a new line, prep the next row
		if b == '\n' {
			row++
			rows = append(rows, []string{})
		}
	}
	return rows
}

func (rr *ResettingReader) Read() ([][]string, error) {
	rr.fh.Seek(0, 0)
	allBytes := []byte{}
	tmpBytes := make([]byte, INITIAL_SLICE)
	for {
		read, err := rr.fh.Read(tmpBytes)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if read > 0 {
			allBytes = append(allBytes, tmpBytes[0:read]...)
		} else {
			return tokenize(allBytes), nil
		}
	}
}

func (rr *ResettingReader) Close() error {
	return rr.fh.Close()
}

func Open(filename string) (*ResettingReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &ResettingReader{file}, nil
}
