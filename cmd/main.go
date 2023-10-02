package main

import (
	"database/sql"
	"fmt"
	"github.com/PGabrielDev/ports_and_adapters_go/internal/adapters"
	"github.com/PGabrielDev/ports_and_adapters_go/internal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productAdapter := adapters.NewProductDb(db)

	productService := application.NewProductService(productAdapter)

	product, _ := productService.Get("f7fe1a5a-ae20-48a4-8a48-cec98bf965e9")
	fmt.Println(product)

}
