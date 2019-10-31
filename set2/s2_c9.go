package set2

import (
	"errors"
)

// FixedBlockPadding pads the input to the outputSize
func FixedBlockPadding(block []byte, outputSize int) ([]byte, error) {
	if len(block) > outputSize {
		return nil, errors.New("Input block is longer than output")
	}

	output := make([]byte, outputSize)
	copy(output[0:len(block)], block)
	for a := len(block); a < outputSize; a++ {
		output[a] = []byte("\x04")[0]
	}

	return output, nil
}
