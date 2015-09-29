package reader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"

	"github.com/rameshvarun/ups/common"
)

// ReadUPS takes a file object and reads it into a PatchData data structure.
func ReadUPS(data []byte) (*common.PatchData, error) {
	reader := bytes.NewReader(data)

	// Verify Patch through Checksum
	patchSumExpected := crc32.ChecksumIEEE(data[:len(data)-4])
	var patchSumActual uint32
	err := binary.Read(bytes.NewBuffer(data[len(data)-4:]), binary.LittleEndian, &patchSumActual)
	if err != nil {
		return nil, err
	}
	if patchSumExpected != patchSumActual {
		return nil, errors.New("Patch failed checksum.")
	}

	// Read and validate the signature.
	signature := make([]byte, 4)
	_, err = io.ReadAtLeast(reader, signature, 4)
	if err != nil {
		return nil, err
	}
	if !common.ValidateSignature(signature) {
		return nil, errors.New("File did not have valid UPS signature.")
	}

	// Read the input and output file sizes.
	inputFileSize, err := ReadVariableLengthInteger(reader)
	if err != nil {
		return nil, err
	}

	// Read the input and output file sizes.
	outputFileSize, err := ReadVariableLengthInteger(reader)
	if err != nil {
		return nil, err
	}

	// Read in the Patch Blocks
	var patchBlocks []common.PatchBlock
	for reader.Len() > 12 {
		// First, read in the offset from the previous block (or beginning of file)
		// as a variable-length integer.
		relativeOffset, err := ReadVariableLengthInteger(reader)
		if err != nil {
			return nil, err
		}

		// Read in the NULL-terminated data portion.
		var data []byte
		for {
			b, err := reader.ReadByte()
			if err != nil {
				return nil, err
			}
			if b == 0 {
				break
			}
			data = append(data, b)
		}

		// Add this to the Patch blocks splice.
		patchBlocks = append(patchBlocks, common.PatchBlock{
			RelativeOffset: relativeOffset,
			Data:           data,
		})
	}

	// Read the checksum of the input file.
	var inputChecksum uint32
	err = binary.Read(reader, binary.LittleEndian, &inputChecksum)
	if err != nil {
		return nil, err
	}

	// Read in the checksum of the output file.
	var outputChecksum uint32
	err = binary.Read(reader, binary.LittleEndian, &outputChecksum)
	if err != nil {
		return nil, err
	}

	// What's left should just be the patch checksum.
	if reader.Len() != 4 {
		return nil, errors.New("File was longer than expected.")
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
