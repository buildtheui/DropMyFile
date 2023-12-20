package network

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"

	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/buildtheui/DropMyFile/pkg/models"
	"github.com/mdp/qrterminal"
)

var myIp string

func PrintLanServerIpQr() {
	serverAddr := GetServerAddr("/")
	qrterminal.Generate(serverAddr, qrterminal.L, os.Stdout)
	fmt.Println("Or go to: " + serverAddr)
}

func getLocalIp() (string, error) {
	connection, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		return "", err
	}

	defer connection.Close()

	localAddr := connection.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

func GetServerPort() string {
	return strconv.Itoa(global.ServerPort)
}

func GetServerAddr(path string) string {
	if myIp == "" {
		var err error
		myIp, err = getLocalIp()

		if err != nil {
			log.Fatal("LAN address could not be found")
		}
	}

	if global.GSession == "" {
		return BuildUrl(models.URL{
			Scheme: "http",
			Host: myIp,
			Port: GetServerPort(),
			Path: path,
			Queries: map[string]string{},
		});
	}

	return BuildUrl(models.URL{
		Scheme: "http",
		Host: myIp,
		Port: GetServerPort(),
		Path: path,
		Queries: map[string]string{
			"s": global.GSession,
		},
	}) 
}

func BuildUrl(urlValues models.URL) string {
	rawQuery := ""

	for key, value := range urlValues.Queries {
		rawQuery += fmt.Sprintf("%s=%s&", url.QueryEscape(key), url.QueryEscape(value))
	}

	if rawQuery != "" {
		rawQuery = rawQuery[:len(rawQuery)-1]
	}

	host := urlValues.Host

	if urlValues.Port != "" {
		host = fmt.Sprintf("%s:%s", host, urlValues.Port)
	}

	newUrl := &url.URL{
		Scheme:   urlValues.Scheme,
		Host:     host,
		Path:     urlValues.Path,
		RawQuery: rawQuery,
	}

	return newUrl.String()
}