package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// The main function opens a connection to the database, sets up the routes and starts the server
func main() {
	db, err := ConnectToDB()
	if err != nil {
		return
	}
	defer db.Close()

	router := gin.Default()

	router.POST("/products", func(c *gin.Context) {
		CreateProductHandler(db, c)
	})
	router.GET("/products", func(c *gin.Context) {
		SearchProductsHandler(db, c)
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
