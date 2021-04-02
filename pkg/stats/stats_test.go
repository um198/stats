package stats

import (
	// "fmt"
	"reflect"
	"testing"

	"github.com/um198/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 1_000_00, Category: "auto", Status: types.StatusOk},
		{ID: 2, Amount: 1_000_00, Category: "auto", Status: types.StatusOk},
		{ID: 3, Amount: 2_000_00, Category: "mobile", Status: types.StatusOk},
		{ID: 4, Amount: 2_000_00, Category: "mobile", Status: types.StatusOk},
		{ID: 5, Amount: 3_000_00, Category: "it", Status: types.StatusOk},
		{ID: 6, Amount: 3_000_00, Category: "it", Status: types.StatusOk},
		{ID: 7, Amount: 5_000_00, Category: "fun", Status: types.StatusOk},
		{ID: 8, Amount: 5_000_00, Category: "fun", Status: types.StatusOk},
	}

	expected := map[types.Category]types.Money{
		"auto":   1_000_00,
		"mobile": 2_000_00,
		"it":     3_000_00,
		"fun":    5_000_00,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected %v, actual %v", expected, result)
	}

}

func TestPeriodsDynamic(t *testing.T) {
	f := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	s := map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}

	result := PeriodsDynamic(f, s)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected %v, actual %v", expected, result)
	}

}
