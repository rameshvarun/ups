package writer

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
