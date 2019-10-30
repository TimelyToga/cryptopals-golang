package set1

import "crypto/aes"

// DecryptAES128InECB is convoluted but is essentially using AES-128 in
// Electronic Cookbook (ECB) mode.
func DecryptAES128InECB(cipherText []byte, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)

	// bytes
	blockSize := 16
	numBlocks := len(cipherText) / blockSize

	decryptedBlock := make([]byte, len(cipherText))
	for block := 0; block < numBlocks; block++ {
		blockStart := block * blockSize
		blockEnd := (block + 1) * blockSize
		cipher.Decrypt(
			decryptedBlock[blockStart:blockEnd],
			cipherText[blockStart:blockEnd],
		)
	}
	return decryptedBlock
}
