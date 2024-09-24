package persistence

import "github.com/mauriciomd/expense-tracker/types"

type ExpensePersister interface {
	Add(e *types.Expense) error
	Delete(data *types.Expense) error
	ReadAll() ([]*types.Expense, error)
	Update(e *types.Expense) error
}
