package global

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var ServerPort int
var SessionLength int
var GSession string
var TransferFolder string
var AllowPaths = []string{"/assets/"}
var ExcludedFiles = []string{".DS_Store", "Thumbs.db", "desktop.ini"}

// The precedence is ENV variables > CLI flags
func Init() {
	godotenv.Load(".env")

	setGlobalSessionLength()
	setGlobalSession()
	setTransferFolder()
	setServerPort()
}

func setGlobalSessionLength() {
	sessionLen := os.Getenv("DMF_SESSION_LENGTH")

	if sessionLen == "" {
		// sessionLen empty we default to the CLI flag
		return
	}
	
	if sessionLen == "0" {
		// if 0 means the global session protections was deactivated
		// by setting DMF_SESSION_LENGTH=0
		SessionLength = 0
		return
	} 

	num, err := strconv.Atoi(sessionLen)
	if err != nil {
		fmt.Println("Wrong DMF_SESSION_LENGTH defaulting to cli value for session length")
		return
	}

	if num < 0 {
		fmt.Println("DMF_SESSION_LENGTH must be greater than 0 defaulting to cli value for session length")
		return
	}

	SessionLength = num
}

func setGlobalSession() {
	if SessionLength == 0 {
		GSession = ""
		return
	}

	GSession = utils.GenerateRandomString(SessionLength)
}

func setServerPort() {
	port := os.Getenv("DMF_PORT")
	if port != "" {
		num, err := strconv.Atoi(port)
		if err != nil {
			fmt.Println("Wrong DMF_PORT defaulting to cli value for the port")
			return
		}

		if num < 0 {
			fmt.Println("DMF_PORT must be greater than 0 defaulting to cli value for the port")
			return
		}

		ServerPort = num
		return
	}
}

func IsvalidSession(session string) bool {
	// If GSession is empty string means the global session protections was deactivated
	// by setting DMF_SESSION_LENGTH=0 therefore we return this as is a valid session
	return GSession == "" || GSession == session
}

func ValidateSession(c *fiber.Ctx) error {
	path := c.Path()

	for _, allowPath :=range(AllowPaths) {
		if strings.HasPrefix(path, allowPath) {
			return c.Next();
		}
	}

	if !IsvalidSession(c.Query("s")) {
		return c.Status(fiber.StatusUnauthorized).SendString("Access denied")
	}

	return c.Next();
}

func setTransferFolder() {
	folderFromEnv := os.Getenv("DMF_TRANSFER_FOLDER")

	if folderFromEnv == "" && TransferFolder != "" {
		// Default to CLI flag
		return
	}

	if folderFromEnv != "" {
		if _, err := os.Stat(folderFromEnv); os.IsNotExist(err) {
			fmt.Println("a valid DMF_TRANSFER_FOLDER is required")
			os.Exit(1)
		}
		TransferFolder = folderFromEnv 
		return
	}
	
	folder, err := makeDefaultTransferFolder()
	if err != nil {
		fmt.Println("Could not get the transfer files folder path")
		os.Exit(1)
	}
	
	TransferFolder = folder
}

func makeDefaultTransferFolder() (string, error) {
	homeDir, errHomeDir := os.UserHomeDir()
	if errHomeDir != nil {
		fmt.Println("Error getting user's home directory:", errHomeDir)
		return "", nil
	}

	folderName := "TransferedFiles"

	// Create the full path to the Desktop folder
	desktopPath := filepath.Join(homeDir, "Desktop")

	// Create the full path to the target folder on the Desktop
	targetFolderPath := filepath.Join(desktopPath, folderName)

	if _, err := os.Stat(targetFolderPath); os.IsNotExist(err) {
		err := os.MkdirAll(targetFolderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return "", err
		}
	} else if err != nil {
		// Error checking folder existence
		fmt.Println("Error finding default folder:", err)
		return "", err
	}

	return targetFolderPath, nil
}