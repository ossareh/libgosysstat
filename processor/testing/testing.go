package testing

import "os"

type TestHarness struct {
	fh *os.File
}

func MakeTestHarness(filename string) (*TestHarness, error) {
	if fh, err := os.Open(filename); err != nil {
		return nil, err
	} else {
		return &TestHarness{fh}, nil
	}
}

func (th *TestHarness) ReplaceFileHandle(filename string) error {
	if fh, err := os.Open(filename); err != nil {
		return err
	} else {
		th.fh = fh
		return nil
	}
}

func (th *TestHarness) Close() error {
	return th.fh.Close()
}

func (th *TestHarness) Read(p []byte) (int, error) {
	return th.fh.Read(p)
}

func (th *TestHarness) Seek(offset int64, whence int) (int64, error) {
	return th.fh.Seek(offset, whence)
}
