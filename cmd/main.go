package main

import (
	"fmt"
	"generador-dominios/config"
	"generador-dominios/core/usecases"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, usando variables del sistema")
	}

	err = config.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}

	err = usecases.RenamePCService()
	if err != nil {
		fmt.Println("Failed to rename PC:", err)
	}

}
