package main

import (
	"encoding/base64"
	"errors"
	"fmt"
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

func main() {
	fmt.Println("Just run the tests...")
}
