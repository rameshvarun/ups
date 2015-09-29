package writer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriteVariableLengthInteger(t *testing.T) {
	// Test cases.
	cases := []struct {
		Input    uint64
		Expected []byte
	}{{
		Input:    2,
		Expected: []byte{0x82},
	}, {
		Input:    8388608,
		Expected: []byte{0x00, 0x7f, 0x7e, 0x82},
	}}

	// Iterate through the test cases.
	for _, test := range cases {
		out := WriteVariableLengthInteger(test.Input)
		if !bytes.Equal(test.Expected, out) {
			fmt.Printf("Expected {%x} to be {%x}\n", out, test.Expected)
			t.Fail()
		}
	}

}
