package stats

import (
	"github.com/um198/bank/v2/pkg/types"

	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       0,
			Amount:   1_000_00,
			Category: "",
			
		},
		{
			ID:       11,
			Amount:   1_000_00,
			Category: "",
		},
		{
			ID:       3,
			Amount:   1_000_00,
			Category: "",
		},
	}

	fmt.Println(Avg(payments))
	//Output:100000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       0,
			Amount:   1_000_00,
			Category: "2",
		},
		{
			ID:       11,
			Amount:   1_000_00,
			Category: "2",
		},
		{
			ID:       3,
			Amount:   1_000_00,
			Category: "1",
		},
	}

	fmt.Println(TotalInCategory(payments, "2"))

	//Output:200000

}
