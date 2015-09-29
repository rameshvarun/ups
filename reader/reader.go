package reader

import (
	"bufio"
	"encoding/binary"
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

	var patchBlocks []common.PatchBlock
	for true {
		b, err := bufferedReader.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == 0 {
			break
		}

		relativeOffset, err := ReadVariableLengthInteger(bufferedReader)
		if err != nil {
			return nil, err
		}

		xor, err := bufferedReader.ReadByte()
		if err != nil {
			return nil, err
		}

		patchBlocks = append(patchBlocks, common.PatchBlock{
			RelativeOffset: relativeOffset,
			XOR:            xor,
		})
	}

	var inputChecksum int32
	err = binary.Read(bufferedReader, binary.BigEndian, &inputChecksum)
	if err != nil {
		return nil, err
	}

	var outputChecksum int32
	err = binary.Read(bufferedReader, binary.BigEndian, &outputChecksum)
	if err != nil {
		return nil, err
	}

	var patchChecksum int32
	err = binary.Read(bufferedReader, binary.BigEndian, &patchChecksum)
	if err != nil {
		return nil, err
	}

	return &common.PatchData{
		InputFileSize:  inputFileSize,
		OutputFileSize: outputFileSize,
		PatchBlocks:    patchBlocks,
		InputChecksum:  inputChecksum,
		OutputChecksum: outputChecksum,
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
