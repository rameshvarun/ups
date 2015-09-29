package writer

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"

	"github.com/rameshvarun/ups/common"
)

// WriteUPS writes UPS data to a byte array.
func WriteUPS(data *common.PatchData) []byte {
	var buffer bytes.Buffer

	// Write signature.
	buffer.Write(common.Signature)

	// Write input and output file sizes.
	buffer.Write(WriteVariableLengthInteger(data.InputFileSize))
	buffer.Write(WriteVariableLengthInteger(data.OutputFileSize))

	for _, block := range data.PatchBlocks {
		buffer.Write(WriteVariableLengthInteger(block.RelativeOffset))
		buffer.Write(block.Data)
		buffer.WriteByte(0)
	}

	// Write the input and output checksums.
	binary.Write(&buffer, binary.LittleEndian, data.InputChecksum)
	binary.Write(&buffer, binary.LittleEndian, data.OutputChecksum)

	// Checksum the buffer so far.
	checksum := crc32.ChecksumIEEE(buffer.Bytes())
	binary.Write(&buffer, binary.LittleEndian, checksum)

	return buffer.Bytes()
}

// WriteVariableLengthInteger writes a variable-length encoded integer.
// based off of the source for http://www.romhacking.net/utilities/606/
func WriteVariableLengthInteger(value uint64) []byte {
	var data []byte
	x := value & 0x7f
	value >>= 7

	for value != 0 {
		data = append(data, byte(x))
		value--
		x = value & 0x7f
		value >>= 7
	}
	data = append(data, byte(0x80|x))
	return data
}
