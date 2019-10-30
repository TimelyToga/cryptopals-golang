package set1

import (
	"fmt"
	"math"
	"sort"
)

// KeyCandidate is a temp data struct for wrapping the keysize and HammingDistance together
type KeyCandidate struct {
	KeySize         int
	HammingDistance float64
}

// FindKeySize is a nifty function to find the ideal key size for a
// Repeating XOR cipher
func FindKeySize(fileBytes []byte) []int {
	var results []KeyCandidate
	for keySize := 2; keySize <= 40; keySize++ {
		var totalHD float64
		var iterations = 0

		numBlocksToCheck := 4

		// Iterate over combinations of first numBlocksToCheck blocks to find a better approximation for avg HD
		for x := 0; x < numBlocksToCheck; x++ {
			for y := x; y < numBlocksToCheck; y++ {
				if x == y {
					continue
				}

				firstKS := fileBytes[x*keySize : (x+1)*keySize]
				secondKS := fileBytes[y*keySize : (y+1)*keySize]
				totalHD += float64(HammingDistance(firstKS, secondKS))
				iterations++
			}
		}
		finalHD := totalHD / (float64(keySize) * float64(iterations))

		results = append(results, KeyCandidate{KeySize: keySize, HammingDistance: finalHD})
	}

	// Sort by keySize
	sort.Slice(results, func(i, j int) bool {
		return results[i].HammingDistance < results[j].HammingDistance
	})

	// Output top min k keysizes by hamming distance
	var outputKeySizes = make([]int, 3)
	for a := 0; a < len(outputKeySizes); a++ {
		outputKeySizes[a] = results[a].KeySize
		fmt.Printf("The minimum KS is %d and the HD is %f\n", results[a].KeySize, results[a].HammingDistance)
	}

	return outputKeySizes
}

// BreakRepeatingKeyXor searches for the best key size, then will use a tricky method to
// find each byte of the key, then decrypts the message and returns it
// OUTPUT => (decryptedString, keyString)
func BreakRepeatingKeyXor(fileBytes []byte) (string, string) {
	// Search for the keysize in [2, 40]
	keySizeCandidates := FindKeySize(fileBytes)

	// Iterate over the possible keysizes, and find the most reasonable one
	var minScore = math.MaxFloat64
	var bestText string
	var bestKey string
	for _, minKeySize := range keySizeCandidates {
		// Step 5: Make transposed KEYSIZE blocks of KEYSIZE length
		// init 2D array
		numBlocks := len(fileBytes) / minKeySize
		transposedBlocks := make([][]byte, minKeySize)
		for idx := range transposedBlocks {
			transposedBlocks[idx] = make([]byte, numBlocks)
		}

		// Traspose the blocks
		for y := 0; y < minKeySize; y++ {
			for x := 0; x < numBlocks; x++ {
				transposedBlocks[y][x] = fileBytes[x*minKeySize+y]
			}
		}

		// Now find XOR byte for each block
		var keyBytes = make([]byte, minKeySize)
		for idx := range keyBytes {
			_, curKeyByte, _ := FindSingleByteXOR(transposedBlocks[idx])
			keyBytes[idx] = curKeyByte
		}

		// fmt.Println(string(keyBytes))

		// Decrypt the message
		output := string(RepeatedXor(fileBytes, keyBytes))
		curTextScore := ScoreEnglishText(output)
		if curTextScore < minScore {
			minScore = curTextScore
			bestText = output
			bestKey = string(keyBytes)
		}
	}

	return bestText, bestKey
}
