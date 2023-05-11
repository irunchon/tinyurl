package shortening

import "math/rand"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const shortURLLength = 10

func ShorteningURL() string {
	//for {
	shortURL := make([]rune, shortURLLength)
	for i := 0; i < shortURLLength; i++ {
		shortURL[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
	}
	// TO DO: check repetitions in storage
	//}
	return string(shortURL)
}
