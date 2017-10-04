package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"log"
)

// Product struct provides the data model for productstore
type Product struct {
	ID          int
	Title       string
	Description string
	Price       float32
}

var db *sql.DB
func init() {
	var err error

	db, err = sql.Open("postgres", "postgres://user:pass@localhost/productstore")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	product := Product{
		Title:       "Amazon Echo",
		Description: "Amazon Echo - Black",
		Price:       179.99,
	}

	// Insert a product
	//
	createProduct(product)

	// Read all product records
	//
	getProducts()
}

// createProduct inserts product values into product table
func createProduct(prd Product) {

	result, err := db.Exec(
		"INSERT INTO products(title, description, price) VALUES($1,$2, $3)",
			prd.Title, prd.Description, prd.Price)

	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()

	fmt.Printf("Product with id=%d created successfully (%d row affected)\n",
		lastInsertID, rowsAffected)
}

// getProducts reads all records from the product table
func getProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}

	defer rows.Close()
	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, prd)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, pr := range products {
		fmt.Printf("%s, %s, $%.2f\n", pr.Title, pr.Description, pr.Price)
	}
}
