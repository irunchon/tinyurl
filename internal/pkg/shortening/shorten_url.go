package shortening

import "math/rand"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const ShortURLLength = 10

func GenerateURL() string {
	shortURL := make([]rune, ShortURLLength)
	for i := 0; i < ShortURLLength; i++ {
		shortURL[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
	}
	return string(shortURL)
}
