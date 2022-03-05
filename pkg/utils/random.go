package utils

import (
	"math/rand"
	"os"
)

func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func WriteConfig(config map[string]string, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	for k, v := range config {
		_, err := f.WriteString(k + ": " + v + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
