package writer

import (
	"bytes"
	"testing"
)

func TestWriteVariableLengthInteger(t *testing.T) {
	data := WriteVariableLengthInteger(8388608)
	if !bytes.Equal(data, []byte{0x00, 0x7f, 0x7e, 0x82}) {
		t.Fail()
	}
}
