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
	var lastDotIdx int

	if lastDotIdx = strings.LastIndex(filename, "."); lastDotIdx == -1 {
		return filename, ""
	}

	ext := filename[lastDotIdx:]

	return strings.TrimSuffix(filename, ext), ext
}

func RenameFileToUnique(filename string) string {
	basename, ext := SplitFile(filename)

	return fmt.Sprintf("%s-%s%s", basename, GenerateRandomString(4), ext)
}

func FormatHumanDate(dateTime time.Time) string {
	now := time.Now()

	hoursDiff := now.Sub(dateTime).Hours()

	if (hoursDiff < 24) {
		return "Today"
	} else if hoursDiff < 48 {
		return "Yesterday"
	} else {
		// Format the date in Jan 02, 2006 format
		return dateTime.Format("Jan 02, 2006")
	}
}

func FormatSize(size int64) string {
	const (
		B  = 1
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
		PT = 1024 * TB
	)

	switch {
	case size < KB:
		return fmt.Sprintf("%dB", size)
	case size < MB:
		return fmt.Sprintf("%.1fKB", float64(size)/KB)
	case size < GB:
		return fmt.Sprintf("%.1fMB", float64(size)/MB)
	case size < TB:
		return fmt.Sprintf("%.1fGB", float64(size)/GB)
	default:
		return fmt.Sprintf("%.1fTB", float64(size)/TB)
	}
}

// containsString function checks if a string is present in a slice
func ContainsString(slice []string, str string) bool {
    for _, element := range slice {
        if element == str {
            return true
        }
    }
    return false
}