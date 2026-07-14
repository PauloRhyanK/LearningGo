package main

import (
	"fmt"
	"log"

	"drive-manager/internal/auth"
	"drive-manager/internal/driveclient"
	"drive-manager/internal/synclogic"
)

func main() {
	fmt.Println("Iniciando o Drive Manager...")

	// 1. Autenticação
	httpClient, err := auth.GetGoogleClient("./infra/credentials.json")
	if err != nil {
		log.Fatalf("Erro ao autenticar: %v", err)
	}
	fmt.Println("Cliente autenticado com sucesso!")

	// 2. Iniciar Ferramentas
	driverTools, err := driveclient.NewDriveClient(httpClient)
	if err != nil {
		log.Fatalf("Erro ao iniciar ferramentas: %v", err)
	}
	fmt.Println("Ferramentas iniciadas com sucesso!")

	// 3. Iniciar Gerenciador de Sincronização
	syncManager := synclogic.NewManagerSync(driverTools)
	fmt.Println("Gerenciador de sincronização iniciado com sucesso!")

	_ = syncManager

}
