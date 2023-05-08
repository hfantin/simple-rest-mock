package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type environment struct {
	ServerPort      string
	TargetServer    string
	WriteFile       bool
	UseHTTPS        bool
	CertificatePath string
	KeyPath         string
	Endpoints       []string
}

var Env *environment

func init() {
	Env = &environment{"9000", "", false, false, "", "", make([]string, 0)}
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using default values....")
	}
	Env.ServerPort = os.Getenv("SERVER_PORT")
	Env.TargetServer = os.Getenv("TARGET_SERVER")
	Env.WriteFile, _ = strconv.ParseBool(os.Getenv("WRITE_FILE"))
	Env.UseHTTPS, _ = strconv.ParseBool(os.Getenv("USE_HTTPS"))
	Env.CertificatePath = os.Getenv("CERTIFICATE_PATH")
	Env.KeyPath = os.Getenv("KEY_PATH")
	endpoints := os.Getenv("ENDPOINTS")

	if len(endpoints) > 0 {
		Env.Endpoints = strings.FieldsFunc(endpoints, func(c rune) bool {
			return c == ';'
		})
	}
	showBanner(versionNumber)
	fmt.Printf("- server port: %s\n", Env.ServerPort)
	fmt.Printf("- targetServer: %s\n", Env.TargetServer)
	fmt.Printf("- write file: %t\n", Env.WriteFile)
	fmt.Printf("- use https: %t\n", Env.UseHTTPS)
	if len(Env.Endpoints) == 0 {
		fmt.Println("you should intercept at least one endpoint")
		os.Exit(1)
	}
	fmt.Println("- endpoints: ")
	for _, v := range Env.Endpoints {
		fmt.Println("  ", v)
	}
}

func showBanner(version string) {
	banner := `
███████ ██ ███    ███ ██████  ██      ███████     ██████  ███████ ███████ ████████     ███    ███  ██████   ██████ ██   ██
██      ██ ████  ████ ██   ██ ██      ██          ██   ██ ██      ██         ██        ████  ████ ██    ██ ██      ██  ██
███████ ██ ██ ████ ██ ██████  ██      █████       ██████  █████   ███████    ██        ██ ████ ██ ██    ██ ██      █████
     ██ ██ ██  ██  ██ ██      ██      ██          ██   ██ ██           ██    ██        ██  ██  ██ ██    ██ ██      ██  ██
███████ ██ ██      ██ ██      ███████ ███████     ██   ██ ███████ ███████    ██        ██      ██  ██████   ██████ ██   ██
                                                                                                           version %s
`
	fmt.Printf(banner, version)
}
