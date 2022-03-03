package repo

// Database methods is used through this interface
type DatabaseRepo interface {
	GetInfo(productName, customerPhone string,price float64,installmentDuration int) (float64,error)
	Pay(productName, customerPhone string,installmentDuration int, payingAmount float64) (float64,error)
	NewProductInstallment(id,productName, customerPhone string,productPrice, additionPrice float64, productCategory,installmentRange int) error
}
