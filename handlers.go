package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Parses request body and validates input before inserting a new product
func CreateProductHandler(db *sql.DB, c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Empty request payload"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if product.Name == "" || product.Category == "" || product.SKU == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	err := InsertProduct(db, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

// Grabs the query parameter text and passes it in to the search function
func SearchProductsHandler(db *sql.DB, c *gin.Context) {
	queryText := c.Query("q")

	if queryText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search - empty query"})
		return
	}

	products, err := SearchProducts(db, queryText)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No matching products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}
