package common

import "bytes"

type PatchData struct {
	InputFileSize  uint64
	OutputFileSize uint64

	PatchBlocks []PatchBlock

	InputChecksum  int32
	OutputChecksum int32
}

// PatchBlock represents a single change from the original file to the modified
// file.
type PatchBlock struct {
	RelativeOffset uint64
	XOR            byte
}

func ValidateSignature(signature []byte) bool {
	return bytes.Equal(signature, []byte{'U', 'P', 'S', '1'})
}
