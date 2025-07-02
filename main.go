package main

import (
	"bulk-file-converter/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}
}
