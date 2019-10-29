package set1

import (
	"encoding/hex"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex2Base__01_01(t *testing.T) {
	hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexBytes, err := hex.DecodeString(hexInput)
	assert.Nil(t, err)

	base64Output := HexToBase64(hexBytes)
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	assert.Equal(t, expectedOutput, base64Output)
}

func TestXOR__01_02(t *testing.T) {
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

func TestFindMostLikelyEnglish__01_03(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputBytes, err := hex.DecodeString(input)
	assert.Nil(t, err)

	minMsg, minByte, _ := FindSingleByteXOR(inputBytes)
	fmt.Println(minMsg)

	assert.Equal(t, uint8(0x58), minByte)
	assert.Equal(t, "Cooking MC's like a pound of bacon", minMsg)
}

func TestFindSingleCharacterXOR__01_04(t *testing.T) {
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

func TestRepeatingXor(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	inputBytes := []byte(input)
	fmt.Println(inputBytes)

	key := "ICE"
	keyBytes := []byte(key)

	expectedOutput := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	outputArray := RepeatedXor(inputBytes, keyBytes)

	assert.Equal(t, expectedOutput, hex.EncodeToString(outputArray))

	// TODO: Ok. I'm close here
	/* But I can't figure out what's going wrong. Somehow the output also has a newline in it??
	 */
}
