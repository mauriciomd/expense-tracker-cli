package utils

import "github.com/mauriciomd/expense-tracker/types"

func Filter(collection []*types.Expense, f func(e *types.Expense) bool) []*types.Expense {
	output := make([]*types.Expense, 0)

	for _, exp := range collection {
		if f(exp) {
			output = append(output, exp)
		}
	}

	return output
}
