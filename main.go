package main

import (
	"embed"

	"github.com/buildtheui/DropMyFile/pkg/cmd"
	"github.com/buildtheui/DropMyFile/pkg/global"
)

//go:embed views/*
var views embed.FS

//go:embed views/assets/*
var assets embed.FS

func main() {
	// Set views files to global variable so files are embed at compile time
	global.ViewsContent = views
	global.ViewsAssets = assets

	// Init command line
	cmd.Execute()
}