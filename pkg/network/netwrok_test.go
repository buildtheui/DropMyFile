package network

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/buildtheui/DropMyFile/pkg/models"
)

func TestGetServerPort(t *testing.T) {
	global.ServerPort = 3000;

	port := GetServerPort()
	if port != strconv.Itoa(global.ServerPort) {
		t.Errorf("GetServerPort() = %q, want %q", port, strconv.Itoa(global.ServerPort))
	}
}

func TestGetServerAddr(t *testing.T) {
	fetchLocalIP = func() (string, error) {
		return "127.0.0.1", nil
	}

	global.ServerPort = 3000;

	var cases = []struct {
		name     string
		path     string
		want     string
	}{
		{"1", "/", fmt.Sprintf("http://127.0.0.1:%v/", global.ServerPort)},
		{"2", "/test", fmt.Sprintf("http://127.0.0.1:%v/test", global.ServerPort)},
		{"3", "/other-path", fmt.Sprintf("http://127.0.0.1:%v/other-path", global.ServerPort)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GetServerAddr(c.path)
			if got != c.want {
				t.Errorf("GetServerAddr(%q) = %q, want %q", c.path, got, c.want)
			}
		})
	}
}

func TestGetServerAddrWithSession(t *testing.T) {
	fetchLocalIP = func() (string, error) {
		return "127.0.0.1", nil
	}

	global.ServerPort = 3000;
	global.GSession = "1A2b3C";

	var cases = []struct {
		name     string
		path     string
		want     string
	}{
		{"1", "/", fmt.Sprintf("http://127.0.0.1:%v/?s=%s", global.ServerPort, global.GSession)},
		{"2", "/test", fmt.Sprintf("http://127.0.0.1:%v/test?s=%s", global.ServerPort, global.GSession)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GetServerAddr(c.path)
			if got != c.want {
				t.Errorf("GetServerAddr(%q) = %q, want %q", c.path, got, c.want)
			}
		})
	}
}

func TestBuildUrl(t *testing.T) {
	var cases = []struct {
		name     string
		urlValues models.URL
		want     string
	}{
		{"1", models.URL{
			Scheme: "http",
			Host:   "127.0.0.1",
			Port:   "8080",
			Path:   "/",
			Queries: map[string]string{},
		}, "http://127.0.0.1:8080/"},
		{"2", models.URL{
			Scheme: "http",
			Host:   "127.0.0.1",
			Port:   "8080",
			Path:   "/test",
			Queries: map[string]string{
				"foo": "bar",
			},
		}, "http://127.0.0.1:8080/test?foo=bar"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := BuildUrl(c.urlValues)
			if got != c.want {
				t.Errorf("BuildUrl(%v) = %q, want %q", c.urlValues, got, c.want)
			}
		})
	}
}
