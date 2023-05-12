package shortening

import "math/rand"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const shortURLLength = 10

func GenerateURL() string {
	shortURL := make([]rune, shortURLLength)
	for i := 0; i < shortURLLength; i++ {
		shortURL[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
	}
	return string(shortURL)
}
