package reader

import (
	"io"
	"io/ioutil"
)

const INITIAL_ROWS = 1

const (
	SCAN_START = iota // looking for non-whitespace
	SCAN_END
)

type DataSource interface {
	io.ReadSeeker
	io.Closer
}

type ResettingReader struct {
	src DataSource
}

func isWhiteSpace(b byte) bool {
	return (b == '\n' || b == '\r' || b == '\t' || b == ' ')
}

func tokenize(bs []byte) [][]string {
	rows := make([][]string, INITIAL_ROWS)
	var row, startPos int
	rows[row] = []string{}

	state := SCAN_START
	for pos, b := range bs {
		if state == SCAN_START && isWhiteSpace(b) {
			// intentionally left blank
		}

		if state == SCAN_START && !isWhiteSpace(b) {
			startPos = pos
			state = SCAN_END
		}

		if state == SCAN_END && !isWhiteSpace(b) {
			// handle case where this is the last value in the array
			if pos == len(bs)-1 {
				rows[row] = append(rows[row], string(bs[startPos:pos+1]))
			}
		}

		if state == SCAN_END && isWhiteSpace(b) {
			rows[row] = append(rows[row], string(bs[startPos:pos]))
			state = SCAN_START
		}

		if b == '\n' {
			row++
			rows = append(rows, []string{})
		}
	}

	return rows
}

func (rr *ResettingReader) Read() ([][]string, error) {
	rr.src.Seek(0, 0)
	bs, err := ioutil.ReadAll(rr.src)
	if err != nil {
		return nil, err
	}
	return tokenize(bs), nil
}

func (rr *ResettingReader) Close() error {
	return rr.src.Close()
}

func New(src DataSource) *ResettingReader {
	return &ResettingReader{src}
}
