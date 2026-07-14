package driveclient

import (
	"net/http"

	"google.golang.org/api/drive/v3"
)

func NewDriveClient(httpClient *http.Client) (*drive.Service, error) {
	println("Iniciando cliente Drive...")
	return nil, nil
}
