package folderwatch

import (
	"fmt"
	"log"
	"os"

	"github.com/buildtheui/DropMyFile/pkg/models"
	"github.com/buildtheui/DropMyFile/pkg/network"
	"github.com/fsnotify/fsnotify"
)

func WatchFileChanges(folderChange chan<- []string) {
	// Specify the folder to watch
	folderPath := os.Getenv("TRANSFER_FOLDER")

	// Create new watcher.
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

	 // Start listening for events.
	 go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }					
					// Notify folder changed
					folderChange <- []string{event.Name}
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

	// Watches the transfer folder for changes
	err = watcher.Add(folderPath)

	if err != nil {
		log.Fatal(err)
	}

	// Wait forever to keep watching
	select {}
}

func GetTransferFilesInfo() ([]models.FileInfo, error) {
	// Specify the folder to watch
	folderPath := os.Getenv("TRANSFER_FOLDER")

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, err
	}

	// Create a slice to store FileInfo structs
	var fileInfos []models.FileInfo

	// Iterate over files and populate FileInfo structs
	for _, file := range files {
		if !file.IsDir() {
			downloadLink := network.GetServerAddr(fmt.Sprintf("/api/v1/download/%s", file.Name()))

			// Create a FileInfo struct for the current file
			fileInfo := models.FileInfo{
				FileName:    file.Name(),
				DownloadLink: downloadLink,
			}

			// Append FileInfo to the slice
			fileInfos = append(fileInfos, fileInfo)
		}
	}

	return fileInfos, nil
}