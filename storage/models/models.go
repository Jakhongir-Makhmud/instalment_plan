package models


type Pay struct {

	ProductName string `json:"product_name"`
	CustomerPhone string `json:"customer_phone"`
	InstallmentDuration int `json:"installment_duration"`
	PayingAmount float64 `json:"paying_amount"`

}