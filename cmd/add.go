package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/spf13/cobra"
)

const defaultCategory = "misc"

func newAdd(p persistence.ExpensePersister) *cobra.Command {
	var description, category string
	var amount float64

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an expense to the database",
		Run: func(cmd *cobra.Command, args []string) {
			expense, err := addExpense(description, category, amount, p)
			if err != nil {
				fmt.Printf("error: %q\n", err.Error())
				return
			}

			fmt.Printf("Expense added successfully (ID: %d)\n", expense.Id)
		},
	}

	cmd.Flags().StringVarP(&description, "description", "d", "", "the expense name.")
	cmd.Flags().StringVarP(&category, "category", "c", "", "the expense category.")
	cmd.Flags().Float64VarP(&amount, "amount", "a", 0, "the expense amount.")

	return cmd
}

func addExpense(description, category string, amount float64, p persistence.ExpensePersister) (*types.Expense, error) {
	if description == "" {
		return nil, ErrInvalidDescription
	}

	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	category = strings.TrimSpace(category)
	if category == "" {
		category = defaultCategory
	}

	id, err := getNextId(p)
	if err != nil {
		return nil, err
	}

	e := &types.Expense{
		Id:          id,
		Description: description,
		Category:    category,
		Amount:      amount,
		Date:        time.Now().Local(),
	}

	p.Add(e)

	return e, nil
}

func getNextId(p persistence.ExpensePersister) (uint, error) {
	data, err := p.ReadAll()
	if err != nil {
		return 0, err
	}

	max := uint(0)
	for _, expense := range data {
		if expense.Id > max {
			max = expense.Id
		}
	}

	return max + 1, nil
}
