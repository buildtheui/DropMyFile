package models

import "time"

// Constants
const (
	Files_modified = "files_modified"
)

// FileInfo represents information about a file, including its name and download link.
type FileInfo struct {
	File_name     	string `json:"fileName"`
	Size	      	string `json:"size"`
	Mod_at        	string `json:"modTime"`
	Download_link 	string `json:"downloadLink"`
	Mod_at_to_sort 	time.Time `json:"-"`
}

type WSResponse struct {
	Event_name string `json:"eventName"`
	Payload interface{} `json:"payload"`
}

type URL struct {
	Scheme   string
	Host     string
	Port	 string
	Path     string
	Queries  map[string]string
}
