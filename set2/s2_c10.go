package set2

import "fmt"

// CBCEncrypt will encrypt the input in 16 byte blocks
// Currently uses an IV of all ASCII \x00
func CBCEncrypt(plain []byte, key []byte) []byte {
	// In bytes; this may eventually move to the signature
	blockSize := 16

	// Create / fill IV
	initializationVector := make([]byte, blockSize)
	for a := 0; a < blockSize; a++ {
		initializationVector[a] = byte('\x00')
	}

	fmt.Println(initializationVector)

	// Encryp

	return initializationVector
}
