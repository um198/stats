package stats

import (
	"github.com/um198/bank/v2/pkg/types"
	
)

// Avg рассчитывает ср. сумму
func Avg(payments []types.Payment) types.Money {
	var averageAmount types.Money = 0
	i:=0
	for _, payment := range payments {
		if payment.Status!=types.StatusFail {
		averageAmount += payment.Amount
		i++
		}

	}
	return averageAmount/types.Money(i)

}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	var paymentsFromCategory types.Money = 0
	for _, payment := range payments {
		if payment.Category==category{
			if payment.Status!=types.StatusFail {
				paymentsFromCategory += payment.Amount
			}
		} 
	}
	return paymentsFromCategory
}
