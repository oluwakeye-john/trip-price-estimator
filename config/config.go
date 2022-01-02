package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func MustGetEnv(i string) string {
	v := os.Getenv(i)
	if v == "" {
		log.Panicf("Variable %s not found", v)
	}
	return v
}
