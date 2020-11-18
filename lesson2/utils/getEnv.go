package utils

import (
	"log"

	"github.com/joho/godotenv"
)

//GetEnvs returns a map of config var
func GetEnvs(envPath string) map[string]string {
	envs, err := godotenv.Read(envPath)

	if err != nil {
		log.Fatal(err)
	}

	return envs
}
