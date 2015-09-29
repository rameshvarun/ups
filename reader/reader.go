package reader

import (
	"bufio"
	"errors"
	"io"

	"github.com/rameshvarun/ups-tools/common"
)

// ReadUPS takes a file object and reads it into a PatchData data structure.
func ReadUPS(reader io.Reader) (*common.PatchData, error) {
	bufferedReader := bufio.NewReader(reader)

	// Read and validate the signature.
	signature := make([]byte, 4)
	_, err := io.ReadAtLeast(bufferedReader, signature, 4)
	if err != nil {
		return nil, err
	}
	if !common.ValidateSignature(signature) {
		return nil, errors.New("File did not have valid UPS signature.")
	}

	// Read the input and output file sizes.
	inputFileSize, err := ReadVariableLengthInteger(bufferedReader)
	if err != nil {
		return nil, err
	}

	// Read the input and output file sizes.
	outputFileSize, err := ReadVariableLengthInteger(bufferedReader)
	if err != nil {
		return nil, err
	}

	return &common.PatchData{
		InputFileSize:  inputFileSize,
		OutputFileSize: outputFileSize,
		PatchBlocks:    make([]common.PatchBlock, 0),
		InputChecksum:  0,
		OutputChecksum: 0,
	}, nil
}

// ReadVariableLengthInteger reads a variable-length encoded integer.
// Based off of https://github.com/mgba-emu/mgba/blob/31993afd2a9bcadda690248f77d0f62363b82b51/src/util/patch-ups.c#L208
func ReadVariableLengthInteger(reader io.ByteReader) (uint64, error) {
	value := uint64(0)
	shift := uint64(1)

	for true {
		b, err := reader.ReadByte()
		if err != nil {
			return 0, err
		}

		value += (uint64(b&0x7F) * shift)
		if b&0x80 != 0 {
			break
		}
		shift <<= 7
		value += shift
	}
	return value, nil
}
