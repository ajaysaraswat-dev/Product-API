package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ajaysaraswat-dev/ecom/internals/products"
	"github.com/gin-gonic/gin"
)

type Application struct {
	config Config
}

//run - to start the server and listen for incoming requests

func (app *Application) run(h http.Handler) error {
	srv := &http.Server{
		Addr : app.config.addr,
		Handler: h,
		WriteTimeout: time.Second *30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}
	log.Printf("server has started at addr %s",app.config.addr)
	return srv.ListenAndServe()
}
//mount - to register the handler functions for the different routes

func (app *Application) mount() http.Handler {

	//user -> handler GET /products -> Services getProducts -> repo select * FROM PRODUCTS
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message" : "Hello",
		})
	})
	productService := products.NewService() //create a instance of the service layer
	productHandler := products.NewHandler(productService) //pass the service here
	r.GET("/products",productHandler.ListProducts)

	return r
}

type Config struct {
	addr string
	db   dbconfig
}

type dbconfig struct {
	dsn string //dsn -> data source name (it contains db username,pass,host port, dbname)

}