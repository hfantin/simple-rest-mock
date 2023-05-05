package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type environment struct {
	ServerPort   string
	TargetServer string
	WriteFile    bool
}

var Env *environment

func init() {
	Env = &environment{"5000", "", false}
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using default values....")
	}
	Env.ServerPort = os.Getenv("SERVER_PORT")
	Env.TargetServer = os.Getenv("TARGET_SERVER")
	Env.WriteFile, _ = strconv.ParseBool(os.Getenv("WRITE_FILE"))
	log.Println("environment:", Env.ServerPort, Env.TargetServer, Env.WriteFile)

}
