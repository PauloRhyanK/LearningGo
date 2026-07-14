package synclogic

import (
	"google.golang.org/api/drive/v3"
)

func NewManagerSync(driveClient *drive.Service) *syncManager {
	println("Iniciando gerenciador de sincronização...")
	return &syncManager{}
}

type syncManager struct {
	driveClient *drive.Service
}
