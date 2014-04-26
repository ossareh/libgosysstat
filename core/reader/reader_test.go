package reader

import (
	"os"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	bytes := make([][]byte, 5)
	bytes[0] = []byte{'a', 'l', 'i', 'c', 'e'}
	bytes[1] = []byte{'a', 'l', 'i', 'c', 'e', '\n'}
	bytes[2] = []byte{'a', 'l', 'i', 'c', 'e', '\n', ' '}
	bytes[3] = []byte{'a', 'l', 'i', 'c', 'e', ' ', '\n'}
	bytes[4] = []byte{'a', 'l', 'i', 'c', 'e', ' ', ' '}

	expected := []string{"alice"}

	for idx, b := range bytes {
		result := tokenize(b)[0]
		if !reflect.DeepEqual(expected, result) {
			t.Fatalf("Expected %s got %s, element %d", expected, result, idx)
		}
	}
}

func TestResettingReader(t *testing.T) {
	file, err := os.Open("./example.sample")
	if err != nil {
		t.Fatalf("Expecting to find sample file", err)
	}
	rr := NewResettingReader(file)
	defer rr.Close()
	known := [][]string{
		[]string{"foo", "bar", "baz"},
		[]string{"bob"},
		[]string{"annie"},
	}
	firstRun, _ := rr.Read()
	secondRun, _ := rr.Read()

	if !reflect.DeepEqual(known, firstRun) ||
		!reflect.DeepEqual(known, secondRun) {
		t.Fatalf("Expected results to be the same", known, firstRun, secondRun)
	}
}
