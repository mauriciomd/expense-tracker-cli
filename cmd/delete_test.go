package cmd

import (
	"testing"
	"time"

	"github.com/mauriciomd/expense-tracker/types"
)

func TestDelete(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		id := uint(0)
		want := ErrInvalidId
		got := deleteExpense(id, &MockPersistence{})

		if got != want {
			t.Fatalf("got %q want %q", got.Error(), want.Error())
		}
	})

	t.Run("expense not found", func(t *testing.T) {
		mock := &MockPersistence{
			data: []*types.Expense{
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

	t.Run("delete expense", func(t *testing.T) {
		id := uint(2)
		mock := &MockPersistence{
			data: []*types.Expense{
				{Id: id, Description: "valid description", Amount: 10, Date: time.Now()},
				{Id: id + 1, Description: "another valid description", Amount: 20, Date: time.Now()},
			},
		}
		err := deleteExpense(id, mock)
		if err != nil {
			t.Fatalf("got %q want nil", err.Error())
		}

		if len(mock.data) != 1 {
			t.Fatal("expense was not deleted")
		}
	})
}
