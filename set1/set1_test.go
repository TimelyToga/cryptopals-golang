package set1

import (
	"encoding/hex"
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex2Base__S1_C1(t *testing.T) {
	hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexBytes, err := hex.DecodeString(hexInput)
	assert.Nil(t, err)

	base64Output := HexToBase64(hexBytes)
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	assert.Equal(t, expectedOutput, base64Output)
}

func TestXOR__S1_C2(t *testing.T) {
	first := "1c0111001f010100061a024b53535009181c"
	second := "686974207468652062756c6c277320657965"
	expectedOutput := "746865206b696420646f6e277420706c6179"

	firstBytes, err := hex.DecodeString(first)
	assert.Nil(t, err)

	secondBytes, err := hex.DecodeString(second)
	assert.Nil(t, err)

	expectedOutputBytes, err := hex.DecodeString(expectedOutput)
	assert.Nil(t, err)

	output, err := XorBuffer(firstBytes, secondBytes)
	assert.Nil(t, err)

	assert.Equal(t, expectedOutputBytes, output)
}

func TestFindMostLikelyEnglish__S1_C3(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputBytes, err := hex.DecodeString(input)
	assert.Nil(t, err)

	minMsg, minByte, _ := FindSingleByteXOR(inputBytes)
	fmt.Println(minMsg)

	assert.Equal(t, uint8(0x58), minByte)
	assert.Equal(t, "Cooking MC's like a pound of bacon", minMsg)
}

func TestFindSingleCharacterXOR__S1_C4(t *testing.T) {
	lines, err := ReadLines("4.txt")
	assert.Nil(t, err)

	var minMsg string
	var minScore = math.MaxFloat64

	for _, line := range lines {
		lineBytes, err := hex.DecodeString(line)
		assert.Nil(t, err)

		msg, _, score := FindSingleByteXOR(lineBytes)
		if score < minScore {
			minMsg = msg
			minScore = score
		}
	}

	fmt.Println(minMsg)
	assert.Equal(t, "Now that the party is jumping\n", minMsg)
}

func TestRepeatingXor__S1_C5(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	inputBytes := []byte(input)

	key := "ICE"
	keyBytes := []byte(key)

	expectedOutput := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	outputArray := RepeatedXor(inputBytes, keyBytes)

	assert.Equal(t, expectedOutput, hex.EncodeToString(outputArray))
}

func TestEditDistance__S1_C6(t *testing.T) {
	first := "this is a test"
	second := "wokka wokka!!!"

	editDistance := HammingDistance([]byte(first), []byte(second))
	assert.Equal(t, 37, editDistance)
}

// KeyCandidate is a temp data struct for wrapping the keysize and HammingDistance together
type KeyCandidate struct {
	KeySize         int
	HammingDistance float64
}

func FindKeySize(fileBytes []byte) []int {
	// var minKeySize int
	// var minHD = math.MaxFloat64

	var results []KeyCandidate
	for keySize := 2; keySize <= 40; keySize++ {
		// firstKS := fileBytes[0:keySize]
		// secondKS := fileBytes[keySize : 2*keySize]
		// finalHD := float64(HammingDistance(firstKS, secondKS)) / float64(keySize)

		// Iterate over combinations of first 4 blocks to find a better approximation for avg HD
		var totalHD float64
		var iterations = 0
		for x := 0; x < 4; x++ {
			for y := x; y < 4; y++ {
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

		// fmt.Printf("KS: %d\tHD: %f\n", keySize, finalHD)
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

func TestBreakRepeatingKeyXor__S1_C6b(t *testing.T) {
	fileBytes, err := ReadLineBytes("6.txt")
	t.Log(err)
	assert.Nil(t, err)

	// Convert from base64

	// Search for the keysize in [2, 40]
	keySizeCandidates := FindKeySize(fileBytes)

	// Iterate over the possible keysizes, and find the most reasonable one
	var maxScore = math.MaxFloat64
	var bestText string
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

		fmt.Println(string(keyBytes))

		// Decrypt the message
		output := string(RepeatedXor(fileBytes, keyBytes))
		curTextScore := ScoreEnglishText(output)
		if curTextScore < maxScore {
			maxScore = curTextScore
			bestText = output
		}
	}

	fmt.Println("FINAL OUTPUT STRING:")
	fmt.Println(bestText)
}
