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

func (d Database) NewProductInstallment(id,productName, customerPhone string,productPrice, additionPrice float64, productCategory,installmentRange int) error {
	query := `insert into product_installment (id, product_name,customer_phone, categoty_id, product_price, addition_price, installment_range) values ($1,$2,$3,$4,$5,$6,$7)`

	_,err := d.db.Exec(query,id,productName,customerPhone,productCategory,productPrice,additionPrice,installmentRange)
	if err != nil {
		return err
	}
	return nil
}

// this method returns remaining amount of debt and error, if there is error it returns -1
func (d Database) Pay(productName, customerPhone string,installmentDuration int, payingAmount float64) (float64,error) {
	var paidPart, originalPrice float64
	query := `update product_installment set paid_part = paid_part + $1 where product_name = $2, customer_phone = $3, installment_range = $4 returning paid_part, product_price`

	err := d.db.QueryRow(query,payingAmount,productName,customerPhone,installmentDuration).Scan(&paidPart,originalPrice)
	if err != nil {
		return -1, err
	}

	return originalPrice - paidPart,nil
}