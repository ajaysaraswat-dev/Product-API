package main

import (
	"log"
	"os"
)

func main() {
	cfg := Config{
		addr: ":8080",
		db: dbconfig{
			dsn: "",
		},
	}
	api := Application{
		config: cfg,
	}
	h := api.mount() //handler
	if err := api.run(h); err != nil {
		log.Printf("server is failed to started with err: %v", err)
		os.Exit(1)
	}
}