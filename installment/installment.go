package installment

import (
	"errors"
	"math"
)

func CalculateInstallment(category, duration int, productPrice float64) (float64, error) {
	if duration > 24 {
		return 0, errors.New("max range of installment is 24 monthes")
	}

	if category > 2 {
		return 0, errors.New("there is only 3 category which statrts from 0")
	}

	mutipler := calculateMultiplier(category, duration)

	percent := installmentPercent[category] * mutipler

	additionPrice := (productPrice / 100) * float64(percent)

	return productPrice + additionPrice, nil
}


func calculateMultiplier(category, duration int) int {

	mutiply := float64(duration) / float64(installmentDurations[category])

	return int(math.Ceil(mutiply))
}
