package main

import (
	"os"
	"test-backend-golang-api/api"

	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	routersInit := api.InitRouter()

	port := ":" + os.Getenv("PORT")

	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", port)

	server.ListenAndServe()
}
