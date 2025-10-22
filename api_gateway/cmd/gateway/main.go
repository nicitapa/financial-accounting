package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	r := router.Setup(cfg)
	log.Printf("API Gateway listening on %s", cfg.Port)
	err := http.ListenAndServe(cfg.Port, r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
