package common

type PatchData struct {
	InputFileSize  uint64
	OutputFileSize uint64

	PatchBlocks []PatchBlock

	InputChecksum  uint32
	OutputChecksum uint32
}

// PatchBlock represents a single change from the original file to the modified
// file.
type PatchBlock struct {
	RelativeOffset uint64
	Data           []byte
}

var Signature = []byte{'U', 'P', 'S', '1'}
