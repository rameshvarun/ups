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

	return nil, nil
}
