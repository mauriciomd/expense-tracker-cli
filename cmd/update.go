package cmd

import (
	"fmt"
	"slices"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/spf13/cobra"
)

func newUpdate(p persistence.ExpensePersister) *cobra.Command {
	var id uint
	var description, category string
	var amount float64

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update and an expense to the database",
		Run: func(cmd *cobra.Command, args []string) {
			expense, err := updateExpense(id, description, category, amount, p)
			if err != nil {
				fmt.Printf("error: %q", err.Error())
				return
			}

			fmt.Printf("Expense updated successfully (ID: %d)\n", expense.Id)
		},
	}

	cmd.Flags().UintVarP(&id, "id", "i", 0, "the expense id - required")
	cmd.Flags().StringVarP(&description, "description", "d", "", "the expense name.")
	cmd.Flags().StringVarP(&category, "category", "c", "", "the expense category.")
	cmd.Flags().Float64VarP(&amount, "amount", "a", 0, "the expense amount.")

	return cmd
}

func updateExpense(id uint, description, category string, amount float64, p persistence.ExpensePersister) (*types.Expense, error) {
	if id < 1 {
		return nil, ErrInvalidId
	}

	expenses, err := p.ReadAll()
	if err != nil {
		return nil, err
	}

	idx := slices.IndexFunc(expenses, func(e *types.Expense) bool {
		return e.Id == id
	})

	if idx < 0 {
		return nil, ErrExpenseNotFound
	}

	expense := expenses[idx]
	if description != "" {
		expense.Description = description
	}

	if amount > 0 {
		expense.Amount = amount
	}

	if category != "" {
		expense.Category = category
	}

	if err := p.Update(expense); err != nil {
		return nil, err
	}

	return expense, nil
}
