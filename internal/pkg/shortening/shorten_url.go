package shortening

import (
	"crypto/md5"
	"math/big"
	"math/rand"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const ShortURLLength = 10

func GenerateHashForURL(longURL string) string { //string {
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
	result := make([]uint64, 0, 1)
	for ; num > 0; num /= base {
		result = append(result, num%base)
	}
	return result
}

// GenerateRandomHashForURL makes random hash independent of long URL (UNUSED)
func GenerateRandomHashForURL() string {
	shortURL := make([]rune, ShortURLLength)
	for i := 0; i < ShortURLLength; i++ {
		shortURL[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
	}
	return string(shortURL)
}
