package utils

import (
	"regexp"
	"testing"
	"time"
)


func TestGenerateRandomString(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name string
		len int
	}{
		{"0", 0},
		{"10", 10},
		{"100", 100},
		{"1000", 1000},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := GenerateRandomString(c.len)
			if len(s) != c.len {
				t.Errorf("GenerateRandomString(%d) = %d, want %d", c.len, len(s), c.len)
			}
		})
	}
}

func TestSplitFile(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name     string
		filename string
		want     string
		wantExt  string
	}{
		{"1", "test.txt", "test", ".txt"},
		{"2", "test.pdf", "test", ".pdf"},
		{"3", "test.mp3", "test", ".mp3"},
		{"4", "test.zip", "test", ".zip"},
		{"4", "without-extension", "without-extension", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, gotExt := SplitFile(c.filename)
			if got != c.want {
				t.Errorf("SplitFile(%q) = (%q, %q), want (%q, %q)", c.filename, got, gotExt, c.want, c.wantExt)
			}
		})
	}
}

func TestRenameFileToUnique(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name     string
		filename string
	}{
		{"1", "test.txt"},
		{"2", "test.pdf"},
		{"3", "test.mp3"},
		{"4", "test.zip"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RenameFileToUnique(c.filename)
			if len(got) == len(c.filename) {
				t.Errorf("RenameFileToUnique(%q) = %q, want different", c.filename, got)
			}
		})
	}
}

func TestFormatHumanDate(t *testing.T) {
	t.Parallel()

	todayPattern := `^Today$`
	yesterdayPattern := `^Yesterday$`
	datePattern := `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s\d{2},\s\d{4}$`


	var cases = []struct {
		name     string
		dateTime time.Time
		pattern string
		want     string
	}{
		{"1", time.Now(), todayPattern, "Today"},
		{"2", time.Now().Add(-23 * time.Hour), todayPattern, "Today"},
		{"2", time.Now().Add(-24 * time.Hour), yesterdayPattern, "Yesterday"},
		{"3", time.Now().Add(-47 * time.Hour), yesterdayPattern, "Yesterday"},
		{"3", time.Now().Add(-48 * time.Hour), datePattern, "MM DD, YYYY"},
		{"4", time.Now().Add(-72 * time.Hour), datePattern, "MM DD, YYYY"},
		{"5", time.Now().Add(-500 * time.Hour), datePattern, "MM DD, YYYY"},
		{"6", time.Now().Add(-1000 * time.Hour), datePattern, "MM DD, YYYY"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatHumanDate(c.dateTime)
			matched, err := regexp.MatchString(c.pattern, got)
			
			if !matched || err != nil {
				t.Errorf("FormatHumanDate(%v) = %q, want %q", c.dateTime, got, c.want)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name     string
		size     int64
		want     string
	}{
		{"1", 0, "0B"},
		{"2", 1, "1B"},
		{"3", 1024, "1.0KB"},
		{"4", 1024 * 1024, "1.0MB"},
		{"5", 1024 * 1024 * 1024, "1.0GB"},
		{"6", 1024 * 1024 * 1024 * 1024, "1.0TB"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatSize(c.size)
			if got != c.want {
				t.Errorf("FormatSize(%d) = %q, want %q", c.size, got, c.want)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name     string
		slice    []string
		str      string
		want     bool
	}{
		{"1", []string{"a", "b", "c"}, "a", true},
		{"2", []string{"a", "b", "c"}, "d", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ContainsString(c.slice, c.str)
			if got != c.want {
				t.Errorf("ContainsString(%v, %q) = %v, want %v", c.slice, c.str, got, c.want)
			}
		})
	}
}