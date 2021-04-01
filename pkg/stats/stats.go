package stats

import (
	
)

// Avg рассчитывает ср. сумму
func Avg(payments []types.Payment) types.Money {
	var averageAmount types.Money = 0
	for _, payment := range payments {
		// if payment.Status!=types.StatusFail {
		averageAmount += payment.Amount
		// }
	}
	return averageAmount/types.Money(len(payments))

}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	var paymentsFromCategory types.Money = 0
	for _, payment := range payments {
		if payment.Category==category{
			// if payment.Status!=types.StatusFail {
				paymentsFromCategory += payment.Amount
			// }
		} 
	}
	return paymentsFromCategory
}
