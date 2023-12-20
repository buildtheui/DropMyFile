package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/buildtheui/DropMyFile/pkg/api"
	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/buildtheui/DropMyFile/pkg/network"
	"github.com/spf13/cobra"
)

var DMFCmd = &cobra.Command{
	Use: "dmf",
	Short: "Starts DropMyFile application",
	Long: `Starts DropMyFile application to transfer files between users over the same LAN.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(global.TransferFolder); global.TransferFolder != "" && os.IsNotExist(err) {
			return fmt.Errorf("a valid --folder-path is required")
		}

		if global.SessionLength < 0 {
			return fmt.Errorf("--session-length must be greater than 0")
		}

		if global.ServerPort < 0 {
			return fmt.Errorf("--port must be greater than 0")
		}
		
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Runs the server app DropMyFile
		appInit()
	},
}

func appInit() {
	// Init global variables
	global.Init()

	// Init router
	var app = api.RouterInit();

	// Print QR code
	network.PrintLanServerIpQr()

	log.Fatal(app.Listen(":" + network.GetServerPort()))
}

func init() {
	DMFCmd.PersistentFlags().IntVarP(&global.ServerPort, "port", "p", 3000, "Port to listen on")
	DMFCmd.PersistentFlags().IntVarP(&global.SessionLength, "session-length", "s", 6, "random str to secures who can access the files, 0 deactivates it")
	DMFCmd.PersistentFlags().StringVarP(&global.TransferFolder, "folder-path", "f", "", "Folder path to transfer files to and download from")
}


func Execute() {	
	if err := DMFCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
