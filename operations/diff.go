package operations

import (
	"hash/crc32"

	"github.com/rameshvarun/ups/common"
)

// Diff takes in a base buffer, a modified buffer, and returns a PatchData object
// that can be used to write to a UPS file.
func Diff(base []byte, modified []byte) *common.PatchData {
	// Cumulative list of blocks that we construct as we scan through the buffers.
	var blocks []common.PatchBlock

	// The end position of the last patch block that we saw.
	lastBlock := uint64(0)
	// The current block that we are constructing. `nil` if there is no current block.
	var currentBlock *struct {
		Data  []byte
		Start uint64
	}

	for pointer := 0; pointer < len(modified); pointer++ {
		// Determine if the byte has been 'modified'
		var different bool
		if pointer >= len(base) {
			// If the output file is larger than the input, and the byte in this extended
			// region is non-zero, then it is 'modified'
			different = modified[pointer] != 0
		} else {
			// Otherwise, simply check if the byte is different.
			different = modified[pointer] != base[pointer]
		}

		if different {
			// If the current byte is modified, but we are not constructing a block,
			// we need to create an in-progress block, then add in the data.
			if currentBlock == nil {
				currentBlock = &struct {
					Data  []byte
					Start uint64
				}{
					Data:  []byte{},
					Start: uint64(pointer),
				}
			}

			// If we are constructing an in-progress block, just add in the data.
			if currentBlock != nil {
				if pointer >= len(base) {
					currentBlock.Data = append(currentBlock.Data, modified[pointer])
				} else {
					currentBlock.Data = append(currentBlock.Data, base[pointer]^modified[pointer])
				}
			}
		} else {
			if currentBlock != nil {
				// This block has ended.
				blocks = append(blocks, common.PatchBlock{
					Data:           currentBlock.Data,
					RelativeOffset: currentBlock.Start - lastBlock,
				})
				currentBlock = nil

				// lastBlock needs to point to the byte after the unmodified byte that ended
				// the block.
				lastBlock = uint64(pointer) + 1
			}
		}
	}

	// If we ended the loop on a block, then we need to end that block.
	if currentBlock != nil {
		blocks = append(blocks, common.PatchBlock{
			Data:           currentBlock.Data,
			RelativeOffset: currentBlock.Start - lastBlock,
		})
	}

	// Return the full patch data structure.
	return &common.PatchData{
		InputFileSize:  uint64(len(base)),
		OutputFileSize: uint64(len(modified)),
		PatchBlocks:    blocks,
		InputChecksum:  crc32.ChecksumIEEE(base),
		OutputChecksum: crc32.ChecksumIEEE(modified),
	}
}
