package shortening

import (
	"crypto/md5"
	"math/big"
)

const (
	// ShortURLLength is hash length for short URL
	ShortURLLength = 10
	// Features of alphabet:
	// - Large (56 chars) and thus very short resulting strings
	// - Proof against offensive words (removed 'U' and 'u')
	// - Unambiguous (removed 'I', 'l', 'O' and '0')
	alphabet = "123456789ABCDEFGHJKLMNPQRSTVWXYZabcdefghijkmnopqrstvwxyz"
)

// Base - quantity of chars in alphabet used for hashing
var base = uint64(len(alphabet))

// GenerateHashForURL ...
func GenerateHashForURL(longURL string) string {
	hashInBytes := md5.Sum([]byte(longURL))
	numberFromHash := new(big.Int).SetBytes(hashInBytes[:]).Uint64()
	return decimalNumToBaseNString(numberFromHash)[:ShortURLLength]
}

func decimalNumToBaseNString(number uint64) string {
	result := ""
	for ; number > 0; number /= base {
		result = string(alphabet[number%base]) + result
	}
	return result
}
