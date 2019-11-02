package set2

import "testing"

func TestCBCEncryption(t *testing.T) {
	input := make([]byte, 32)
	key := []byte("YELLOW SUBMARINE")
	CBCEncrypt(input, key)
}
