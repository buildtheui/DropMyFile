package models

// Constants
const (
	Files_modified = "files_modified"
)

// FileInfo represents information about a file, including its name and download link.
type FileInfo struct {
	FileName     string `json:"fileName"`
	DownloadLink string `json:"downloadLink"`
}

type WSResponse struct {
	Event_name string `json:"eventName"`
	Payload interface{} `json:"payload"`
}