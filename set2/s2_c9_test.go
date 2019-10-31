package set2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedBlockPadding__S2_C9(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	expectedOutput := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	outputBytes, err := FixedBlockPadding(input, 20)
	assert.Nil(t, err)

	assert.Equal(t, expectedOutput, outputBytes)
}
