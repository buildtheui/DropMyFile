package network

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/mdp/qrterminal"
)

func PrintLanServerIpQr() {
	myIp, _ := getLocalIp()
	serverAddr := fmt.Sprintf("http://%s:%s/?s=%s", myIp, GetServerPort(), global.GSession) 
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
	port := os.Getenv("PORT")

	if (port == "") {
		log.Fatal("PORT is not found in the environment")
	}

	return port
}