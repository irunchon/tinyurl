package shortening

import (
	"crypto/md5"
	"math/big"
)

const (
	// ShortURLLength is hash length for short URL
	ShortURLLength = 10
	alphabet       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
)

// GenerateHashForURL ...
func GenerateHashForURL(longURL string) string {
	hashInBytes := md5.Sum([]byte(longURL))
	numberFromHash := new(big.Int).SetBytes(hashInBytes[:]).Uint64()
	return decimalNumToBase63String(numberFromHash)[:ShortURLLength]
}

func decimalNumToBase63String(num uint64) string {
	numbers := decimalNumberConversionToBaseNNumbers(num, 63)
	result := ""
	for i := len(numbers) - 1; i >= 0; i-- {
		result = result + string(alphabet[numbers[i]])
	}
	return result
}

func decimalNumberConversionToBaseNNumbers(num, base uint64) []uint64 {
	if num == 0 || base == 0 {
		return []uint64{0}
	}
	result := make([]uint64, 0, 1)
	for ; num > 0; num /= base {
		result = append(result, num%base)
	}
	return result
}
