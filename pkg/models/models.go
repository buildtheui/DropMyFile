package models

// FileInfo represents information about a file, including its name and download link.
type FileInfo struct {
	FileName    string `json:"fileName"`
	DownloadLink string `json:"downloadLink"`
}