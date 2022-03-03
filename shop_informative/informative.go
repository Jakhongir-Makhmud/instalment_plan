package shopinformative

import "instalment_plan/storage/repo"


type ShopInformative struct {

	ProductName string `json:"product_name"`
	Price float64 `json:"price"`
	CustomerPhone string `json:"customer_phone"`
	InstallmentDuration int `json:"installment_duration"`

}

func (s ShopInformative) GetInfo(repo repo.DatabaseRepo) {



}
