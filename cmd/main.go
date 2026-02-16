package main

import (
	"log"
	"os"

	"github.com/ajaysaraswat-dev/ecom/pkg/database"
)

func main() {
	cfg := Config{
		addr: ":8080",
		db: dbconfig{
			dsn: "mongodb://localhost:27017",
		},
	}


	client,err := database.NewMongoClient(cfg.db.dsn)
	if err != nil{
		log.Printf("connect to failed with the database %v",err)
	}
	api := Application{
		config: cfg,
		db: client,
	}

	h := api.mount() //handler
	if err := api.run(h); err != nil {
		log.Printf("server is failed to started with err: %v", err)
		os.Exit(1)
		//here we can used the log.fatal also
	}
}