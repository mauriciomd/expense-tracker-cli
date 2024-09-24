package persistence

import (
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/mauriciomd/expense-tracker/utils"
)

type MockPersistence struct {
	Data []*types.Expense
}

func (mp *MockPersistence) Add(e *types.Expense) error {
	mp.Data = append(mp.Data, e)
	return nil
}

func (mp *MockPersistence) Delete(e *types.Expense) error {
	mp.Data = utils.Filter(mp.Data, func(expense *types.Expense) bool {
		return e.Id != expense.Id
	})

	return nil
}

func (mp *MockPersistence) ReadAll() ([]*types.Expense, error) {
	return mp.Data, nil
}

func (mp *MockPersistence) Update(e *types.Expense) error {
	mp.Delete(e)
	mp.Add(e)

	return nil
}
