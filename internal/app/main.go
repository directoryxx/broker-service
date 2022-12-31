package main

import (
	"log"
	"os"

	"broker/config"
	"broker/delivery/http"
)

func main() {
	envSource := "SYSTEM"

	if os.Getenv("BYPASS_ENV_FILE") == "" {
		log.Println("[INFO] Load Config")
		config.LoadConfig()
		envSource = "FILE"
	}

	log.Println("[INFO] Loaded Config : " + envSource)

	http.RunWebserver()
}
