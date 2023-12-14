package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// generateRandomString generates a random string of the specified length
func GenerateRandomString(length int) string {
	ranNumber := rand.New(rand.NewSource(time.Now().UnixMicro()))

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[ranNumber.Intn(len(charset))]
	}

	return string(result)
}

// Split the file in basename and extension
func SplitFile(filename string) (string, string) {
	ext := filename[strings.LastIndex(filename, "."):]

	return strings.TrimSuffix(filename, ext), ext
}

func RenameFileToUnique(filename string) string {
	basename, ext := SplitFile(filename)

	return fmt.Sprintf("%s-%s%s", basename, GenerateRandomString(4), ext)
}

