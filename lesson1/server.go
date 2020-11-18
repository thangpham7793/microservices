package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func handleError(rw *http.ResponseWriter, errorMessage string) {
	http.Error(*rw, errorMessage, http.StatusBadRequest)
	return
}

func handleRequest(rw http.ResponseWriter, data []byte) {
	fmt.Fprintf(rw, "Hello %s\n", data)
}

func getEnvConfig(key string) string {
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal(err)
	}

	return envs[key]
}

func registerHandlers() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handleError(&rw, "Error Parsing Body Data!")
			return
		}
		if len(data) == 0 {
			handleError(&rw, "Body cannot be empty!")
			return
		}
		handleRequest(rw, data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye!")
	})
}

// StartServer starts the server (need this to make it public)
func StartServer() {
	PORT := getEnvConfig("PORT")
	registerHandlers()
	http.ListenAndServe(":"+PORT, nil)
}
