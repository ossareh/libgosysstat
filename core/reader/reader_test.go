package reader

import "testing"

func TestOpeningInvalidResettingReader(t *testing.T) {
	_, err := Open("./doesnotexist")
	if err == nil {
		t.Fatalf("Expecting to not file file")
	}
}

func TestResettingReader(t *testing.T) {
	rr, err := Open("./example.sample")
	if err != nil {
		t.Fatalf("Expecting to find sample file", err)
	}
	defer rr.Close()
	known := [][]string{
		[]string{"foo", "bar", "baz"},
		[]string{"bob"},
		[]string{"annie"},
	}
	firstRun, _ := rr.Read()
	secondRun, _ := rr.Read()

	if len(known) != len(firstRun) || len(known) != len(secondRun) {
		t.Fatal("Excepting results to be same length")
	}

	for idx, row := range known {
		for subIdx, col := range row {
			if col != firstRun[idx][subIdx] ||
				col != secondRun[idx][subIdx] {
				t.Fatalf("Expecting results to be the same",
					col,
					firstRun[idx][subIdx],
					secondRun[idx][subIdx])
			}
		}
	}
}
