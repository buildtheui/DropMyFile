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

var GSession = generateGlobalSession()
var AllowPaths = []string{"/assets/"}
var TransferFolder = getTransferFolder()
var ExcludedFiles = []string{".DS_Store", "Thumbs.db", "desktop.ini"}

func generateGlobalSession() string {
	godotenv.Load(".env")

	sessionLen := os.Getenv("DMF_SESSION_LENGTH")

	if sessionLen == "" || sessionLen == "0" {
		// If sessionLen is empty string or 0 means the global session protections was deactivated
		// by setting DMF_SESSION_LENGTH=0 or not defining it at all
		return ""
	} else {
		num, err := strconv.Atoi(sessionLen)
		if err != nil {
			fmt.Println("Wrong DMF_SESSION_LENGTH defaulting to 6 for session length")
			num = 6
		}

		return utils.GenerateRandomString(num)
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

// The precedence is Config file > ENV variable DMF_TRANSFER_FOLDER > default folder
func getTransferFolder() string {
	godotenv.Load(".env")

	var printFolderPath = func (folderPath string) {
		fmt.Printf("Folder for transfering files located at: %s\n", folderPath)
	}
	folderFromEnv := os.Getenv("DMF_TRANSFER_FOLDER")

	fmt.Println(folderFromEnv)

	if folderFromEnv != "" {
		printFolderPath(folderFromEnv)
		return folderFromEnv
	}
	
	folder, err := makeDefaultTransferFolder()
	if err != nil {
		panic("Could not get the transfer files folder path")
	}
	
	printFolderPath(folder)

	return folder
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