package stats

import (
	"github.com/um198/bank/v2/pkg/types"
)

// Avg рассчитывает ср. сумму
func Avg(payments []types.Payment) types.Money {
	var averageAmount types.Money = 0
	i := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			averageAmount += payment.Amount
			i++
		}

	}
	return averageAmount / types.Money(i)

}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var paymentsFromCategory types.Money = 0
	for _, payment := range payments {
		if payment.Category == category {
			if payment.Status != types.StatusFail {
				paymentsFromCategory += payment.Amount
			}
		}
	}
	return paymentsFromCategory
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	catCol := map[types.Category]types.Money{}

	var payment types.Payment
	for _, payment = range payments {
		categories[payment.Category] += payment.Amount
		catCol[payment.Category]++
	}

	for i := range categories {
		categories[i] /= catCol[i]
	}

	return categories

}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for i := range second {
		result[i] = second[i] - first[i]

	}

	return result 

}
