package set1

import "encoding/hex"

// DetectAES128InECB returns a count of the number of repeated blocks
// in the input, which represents how likely it is that this is AES
func DetectAES128InECB(input []byte) int {
	blockSize := 16
	numBlocks := len(input) / blockSize

	// Hex string => count
	counter := make(map[string]int)

	// Fill the counter with block counts
	for a := 0; a < numBlocks; a++ {
		curHexStr := hex.EncodeToString(input[a*blockSize : (a+1)*blockSize])
		counter[curHexStr]++
	}

	// Return number of duplicates
	return numBlocks - len(counter)
}
