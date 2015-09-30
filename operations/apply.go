package operations

import (
	"errors"
	"hash/crc32"

	"github.com/rameshvarun/ups/common"
)

// Apply applies the patch data to the given base.
func Apply(base []byte, patch *common.PatchData, skipCRC bool) ([]byte, error) {
	if uint64(len(base)) != patch.InputFileSize {
		return nil, errors.New("Base file did not have expected size.")
	}
	if !skipCRC && crc32.ChecksumIEEE(base) != patch.InputChecksum {
		return nil, errors.New("Base file did not have expected checksum")
	}

	output := make([]byte, patch.OutputFileSize)
	copy(output, base)

	pointer := 0
	for _, block := range patch.PatchBlocks {
		pointer += int(block.RelativeOffset)

		for _, b := range block.Data {
			if pointer >= len(base) {
				output[pointer] = b
			} else {
				output[pointer] = base[pointer] ^ b
			}
			pointer++
		}

		pointer++
	}

	if !skipCRC && crc32.ChecksumIEEE(output) != patch.OutputChecksum {
		return nil, errors.New("Patch result did not have expected checksum")
	}

	return output, nil
}
