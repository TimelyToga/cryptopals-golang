package set1

import (
	"encoding/base64"
	"errors"
	"math"
	"strings"
)

// HexToBase64 does what you think..
func HexToBase64(hex []byte) string {
	encoder := base64.StdEncoding
	return encoder.EncodeToString(hex)
}

// XorBuffer - XORs two equal length buffers
func XorBuffer(first []byte, second []byte) ([]byte, error) {
	if len(first) != len(second) {
		return nil, errors.New("Buffers are not equal length")
	}

	output := make([]byte, len(first))

	for idx := range first {
		// XOR operation
		output[idx] = first[idx] ^ second[idx]
	}

	return output, nil
}

// XorByte XORs each byte in the buffer with the single input byte
func XorByte(buffer []byte, single byte) []byte {
	output := make([]byte, len(buffer))
	for idx := range buffer {
		output[idx] = buffer[idx] ^ single
	}
	return output
}

// EnglishLetterFrequencies is a letter to frequency map (in %)
var EnglishLetterFrequencies = map[rune]float64{
	' ': 15.00,
	'e': 12.02,
	't': 9.10,
	'a': 8.12,
	'o': 7.68,
	'i': 7.31,
	'n': 6.95,
	's': 6.28,
	'r': 6.02,
	'h': 5.92,
	'd': 4.32,
	'l': 3.98,
	'u': 2.88,
	'c': 2.71,
	'm': 2.61,
	'f': 2.30,
	'y': 2.11,
	'w': 2.09,
	'g': 2.03,
	'p': 1.82,
	'b': 1.49,
	'v': 1.11,
	'k': 0.69,
	'x': 0.17,
	'q': 0.11,
	'j': 0.10,
	'z': 0.07,
}

// ScoreEnglishText runs a Chi-squared test against this text
// to see how well it resembles english text.
//
// 	Lower numbers => closer to english
func ScoreEnglishText(input string) float64 {
	input = strings.ToLower(input)

	// Map characters => count
	var counts = map[rune]int{}

	// Count character frequency
	for _, character := range input {
		counts[character]++
	}

	// Calculate Chi Squared versus expectations
	var chiSquared float64
	for character, frequency := range EnglishLetterFrequencies {
		expectedObservations := frequency * float64(len(input))
		chiSquared += math.Pow(float64(counts[character])-expectedObservations, 2) / expectedObservations
	}
	return chiSquared
}

// FindSingleByteXOR finds the most likely decrypted string by searching for
// the byte that can be XOR'd against the input and produce the most
// "English-like" output
// OUTPUT => (decryptedString, key, score)
func FindSingleByteXOR(input []byte) (string, byte, float64) {
	var minByte byte
	var minMsg string
	var minScore = math.MaxFloat64

	// Iterate through each byte and find most likely one
	for i := 0; i < int(math.Pow(2, 8)); i++ {
		curKey := byte(i)
		decryptedMsgBytes := XorByte(input, curKey)
		msgString := string(decryptedMsgBytes)

		curScore := ScoreEnglishText(msgString)

		if curScore < minScore {
			minScore = curScore
			minMsg = msgString
			minByte = curKey
		}
	}

	return minMsg, minByte, minScore
}

// RepeatedXor first came up in S1P5 and iterates over input,
// XORing with a cycle of bytes from the key
func RepeatedXor(input []byte, key []byte) []byte {
	output := make([]byte, len(input))
	for idx, inputByte := range input {
		curKeyByte := key[idx%(len(key))]
		output[idx] = inputByte ^ curKeyByte
	}
	return output
}

// RepeatedXorLines Split input string on newlines, then apply RepeatedXOR over each line
func RepeatedXorLines(input string, key []byte) [][]byte {
	lines := strings.Split(input, "\n")
	output := make([][]byte, len(lines))

	for idx, line := range lines {
		output[idx] = RepeatedXor([]byte(line), key)
	}
	return output
}

// EditDistance computes the EditDistance for S1C6
func EditDistance(first [], second string) int {
	// Initialize temp array to store edit distances
	var result = make([][]int, len(first)+1)
	for a := range result {
		result[a] = make([]int, len(second)+1)
	}

	// Solve this in the DP way
	for xIndex, x := range result {
		for yIndex := range x {
			if xIndex == 0 {
				// first is empty, fill with second
				result[xIndex][yIndex] = len(second)
			} else if yIndex == 0 {
				// second is empty, fill with first
				result[xIndex][yIndex] = len(first)
				// Previous chars are the same, recurse
			} else if first[xIndex-1] == second[yIndex-1] {
				result[xIndex][yIndex] = result[xIndex-1][yIndex-1]
			} else {
				// Not the same, find out which is the minimum edit distance
				result[xIndex][yIndex] = MinOf(
					result[xIndex][yIndex-1],   // Insert
					result[xIndex-1][yIndex],   // Remove
					result[xIndex-1][yIndex-1], // Replace
				)
			}
		}
	}

	// Return the final edit distance
	return result[len(first)][len(second)]
}
