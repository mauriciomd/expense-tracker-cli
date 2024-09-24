package cmd

import (
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/mauriciomd/expense-tracker/utils"
)

type MockPersistence struct {
	data []*types.Expense
}

func (mp *MockPersistence) Add(e *types.Expense) error {
	mp.data = append(mp.data, e)
	return nil
}

func (mp *MockPersistence) Delete(e *types.Expense) error {
	mp.data = utils.Filter(mp.data, func(expense *types.Expense) bool {
		return e.Id != expense.Id
	})

	return nil
}

func (mp *MockPersistence) ReadAll() ([]*types.Expense, error) {
	return mp.data, nil
}

func (mp *MockPersistence) Update(e *types.Expense) error {
	mp.Delete(e)
	mp.Add(e)

	return nil
}
