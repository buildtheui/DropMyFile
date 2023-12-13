package main

import (
	"log"

	"github.com/buildtheui/DropMyFile/pkg/api"
	"github.com/buildtheui/DropMyFile/pkg/network"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	var app = api.RouterInit();

	network.PrintLanServerIpQr()

	log.Fatal(app.Listen(":" + network.GetServerPort()))
}