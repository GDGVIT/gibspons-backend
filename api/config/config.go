package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfigs() {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../")

	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	loadDBConfig()
	log.Println("[CONFIG] Loaded all configs")
}
