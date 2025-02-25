package main

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error: faild to load .env %v\n", err)
	}
}
