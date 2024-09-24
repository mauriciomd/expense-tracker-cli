package cmd

import (
	"fmt"
	"slices"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/spf13/cobra"
)

func newList(p persistence.ExpensePersister) *cobra.Command {
	var id uint
	var description, category string
	var amount float64

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		Run: func(cmd *cobra.Command, args []string) {
			expenses, err := listExpenses(p)
			if err != nil {
				fmt.Printf("error: %q", err.Error())
				return
			}

			if len(expenses) == 0 {
				fmt.Println("There are no expenses.")
				return
			}

			for _, e := range expenses {
				fmt.Println(e)
			}
		},
	}

	cmd.Flags().UintVarP(&id, "id", "i", 0, "the expense id - required")
	cmd.Flags().StringVarP(&description, "description", "d", "", "the expense name.")
	cmd.Flags().StringVarP(&category, "category", "c", "", "the expense category.")
	cmd.Flags().Float64VarP(&amount, "amount", "a", 0, "the expense amount.")

	return cmd
}

func listExpenses(p persistence.ExpensePersister) ([]*types.Expense, error) {
	expenses, err := p.ReadAll()

	if err != nil {
		return nil, err
	}

	if len(expenses) == 0 {
		return expenses, nil
	}

	slices.SortFunc(expenses, func(a, b *types.Expense) int {
		if a.Date.Equal(b.Date) {
			return 0
		}

		if a.Date.Before(b.Date) {
			return -1
		}

		return 1
	})

	return expenses, nil
}
