package adapters_test

import (
	"database/sql"
	"github.com/PGabrielDev/ports_and_adapters_go/internal/adapters"
	"github.com/PGabrielDev/ports_and_adapters_go/internal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var db *sql.DB

func setUp() {
	db, _ = sql.Open("sqlite3", ":memory:")
	createTable(db)
	createProduct(db)
}

func createTable(db *sql.DB) {
	stmt, err := db.Prepare(`
								CREATE TABLE products (
								  "id" string,
								  "name", string,
								  "price" float,
								  status string  
								);
							`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()

}

func createProduct(db *sql.DB) {
	stmt, err := db.Prepare(`
								INSERT INTO products values ("abc", "lapis", 0, "disable")
							`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()

}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer db.Close()
	producDb := adapters.NewProductDb(db)

	result, err := producDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "lapis", result.GetName())
	require.Equal(t, 0.0, result.GetPrice())
	require.Equal(t, "disable", result.GetStatus())

}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer db.Close()
	producDb := adapters.NewProductDb(db)

	product := application.NewProduct()
	product.Name = "lapis"
	product.Price = 25

	result, err := producDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "lapis", result.GetName())
	require.Equal(t, 25.0, result.GetPrice())
	require.Equal(t, "disable", result.GetStatus())

	product.Status = "enable"
	result, err = producDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "lapis", result.GetName())
	require.Equal(t, 25.0, result.GetPrice())
	require.Equal(t, "enable", result.GetStatus())
}
