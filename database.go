package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	filePath = "./database.db"
)

// Opens a connection to the database and returns a pointer to the sql.DB object and any errors found
func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Inserts a product into the database with the given product information and returns any errors found
func InsertProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("INSERT INTO products (name, category, sku) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Name, product.Category, product.SKU)
	if err != nil {
		return err
	}

	return nil
}

// Executes a SQL query to search for products with the given text string and returns an array of Product structs
func SearchProducts(db *sql.DB, text string) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, name, category, sku, time_created FROM products WHERE name LIKE ? OR category LIKE ? OR sku LIKE ?", "%"+text+"%", "%"+text+"%", "%"+text+"%")
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.SKU, &product.TimeStamp)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}
