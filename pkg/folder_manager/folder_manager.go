package foldermanager

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/buildtheui/DropMyFile/pkg/models"
	"github.com/buildtheui/DropMyFile/pkg/network"
	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/fsnotify/fsnotify"
)

func WatchFileChanges(folderChange chan<- []string) {
	// Specify the folder to watch
	folderPath := global.TransferFolder

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
	folderPath := global.TransferFolder

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

			fileInfo, err := file.Info()
			if err != nil {
				fmt.Println("Error getting file info:", err)
				continue
			}

			mod_at := fileInfo.ModTime()

			// Create a FileInfo struct for the current file
			fileData := models.FileInfo{
				File_name:    file.Name(),
				Size: utils.FormatSize(fileInfo.Size()),
				Mod_at: utils.FormatHumanDate(mod_at),
				Mod_at_to_sort: mod_at,
				Download_link: downloadLink,
			}

			// Append FileInfo to the slice
			fileInfos = append(fileInfos, fileData)
		}
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Mod_at_to_sort.After(fileInfos[j].Mod_at_to_sort)
	})

	var filteredFiles []models.FileInfo

	for _, file := range fileInfos {
		if !utils.ContainsString(global.ExcludedFiles, file.File_name) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles, nil
}