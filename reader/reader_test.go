package reader

import (
	"bytes"
	"testing"
)

func TestReadVariableLengthInteger(t *testing.T) {
	i, err := ReadVariableLengthInteger(bytes.NewReader([]byte{0x00, 0x7f, 0x7e, 0x82}))
	if err != nil {
		t.Error(err)
	}

	if i != 8388608 {
		t.Fail()
	}
}
