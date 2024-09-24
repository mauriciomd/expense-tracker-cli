package cmd

import (
	"testing"
	"time"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
)

func TestUpdateExpense(t *testing.T) {
	id := uint(2)

	mock := &persistence.MockPersistence{
		Data: []*types.Expense{
			{Id: id, Description: "valid description", Amount: 10, Category: "education", Date: time.Now()},
		},
	}

	t.Run("invalid id", func(t *testing.T) {
		id := uint(0)
		description := ""
		category := ""
		amount := 10.0
		want := ErrInvalidId
		_, got := updateExpense(id, description, category, amount, &persistence.MockPersistence{})

		if got != want {
			t.Fatalf("got %q want %q", got.Error(), want.Error())
		}
	})

	t.Run("expense not found", func(t *testing.T) {
		mock := &persistence.MockPersistence{
			Data: []*types.Expense{
				{Id: 2, Description: "valid description", Amount: 10, Date: time.Now()},
			},
		}
		id := uint(1)
		want := ErrExpenseNotFound
		got := deleteExpense(id, mock)

		if got != want {
			t.Fatalf("got %q want %q", got.Error(), want.Error())
		}
	})

	t.Run("change description", func(t *testing.T) {
		description := "new description"
		amount := 0.0
		category := ""

		expense, err := updateExpense(id, description, category, amount, mock)
		assertErrorIsNil(t, err)
		assertAny(t, expense.Description, description)
	})

	t.Run("change amount", func(t *testing.T) {
		description := "breakfast"
		amount := 20.0
		category := "valid category"

		expense, err := updateExpense(id, description, category, amount, mock)
		assertErrorIsNil(t, err)
		assertAny(t, expense.Amount, amount)
	})

	t.Run("change category", func(t *testing.T) {
		description := ""
		amount := 0.0
		category := "education"

		expense, err := updateExpense(id, description, category, amount, mock)
		assertErrorIsNil(t, err)
		assertAny(t, expense.Category, category)
	})

	t.Run("update expense", func(t *testing.T) {
		description := "breakfast"
		category := "food"
		amount := 13.50

		_, err := updateExpense(id, description, category, amount, mock)
		assertErrorIsNil(t, err)

		data := mock.Data[0]
		assertAny(t, data.Description, description)
		assertAny(t, data.Category, category)
		assertAny(t, data.Amount, amount)
	})
}
