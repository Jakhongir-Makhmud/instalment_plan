package shopinformative

import "instalment_plan/storage/repo"


type ShopInformative struct {

	ProductName string `json:"product_name"`
	Price float64 `json:"price"`
	CustomerPhone string `json:"customer_phone"`
	InstallmentDuration int `json:"installment_duration"`

}

// GetInfo returns total amounts paid for a debt
func (s ShopInformative) GetInfo(repo repo.DatabaseRepo) (float64,error) {

	totalPaid ,err := repo.GetInfo(s.ProductName,s.CustomerPhone,s.Price,s.InstallmentDuration)
	if err != nil {
		return -1, err
	}
	return totalPaid,nil

}
