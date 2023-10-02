package adapters

import (
	"database/sql"
	"github.com/PGabrielDev/ports_and_adapters_go/internal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{
		db: db,
	}
}

func (r ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	err := r.db.QueryRow("select id, name, price, status where id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return &product, nil
	}
	return &product, err
}

func (r ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	r.db.QueryRow("select id from products where id = ?", product.GetId()).Scan(&rows)
	if rows > 0 {
		_, err := r.update(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := r.create(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (r ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := r.db.Prepare(`
									"INSERT INTO products ("id", "name", "price", "status") VALUES (?, ?, ?, ?)"
								`)

	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	return product, err
}

func (r ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := r.db.Exec("update products set name = ?, set price = ?, set status = ? where id = ? ",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return nil, err
	}
	return product, err
}
