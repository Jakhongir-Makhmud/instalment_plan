package database

import (
	"github.com/jmoiron/sqlx"
)

// Main struct to interact with database
type Database struct {
	db *sqlx.DB
}

// NewDatabase function initalizes main Database struct
func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (d Database) GetInfo(productName, customerPhone string,price float64,installmentDuration int) (float64,error) {
	var paidPart float64
	query := `select paid_part from product_installment where product_name = $1, customer_phone = $2, product_price = $3, installment_range = $4`
	err := d.db.QueryRow(query,productName,customerPhone,price,installmentDuration).Scan(&paidPart)

	if err != nil {
		return -1, err
	}

	return paidPart,nil

}
