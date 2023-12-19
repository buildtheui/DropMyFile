package global

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var GSession = utils.GenerateRandomString(6)
var AllowPaths = []string{"/assets/"}
var TransferFolder = getTransferFolder()

func IsvalidSession(session string) bool {
	return GSession == session
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