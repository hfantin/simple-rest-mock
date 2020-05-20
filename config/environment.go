package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	ServerPort string
}

var Env *environment

func init() {
	Env = &environment{"5000"}
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using default values....")
	}
	Env.ServerPort = os.Getenv("SERVER_PORT")

}
