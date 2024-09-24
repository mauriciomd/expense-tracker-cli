package cmd

import (
	"fmt"
	"slices"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/spf13/cobra"
)

func newDelete(p persistence.ExpensePersister) *cobra.Command {
	var id uint

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Remove an expense from the database",
		Run: func(cmd *cobra.Command, args []string) {
			err := deleteExpense(id, p)
			if err != nil {
				fmt.Printf("error: %q\n", err.Error())
				return
			}

			fmt.Println("Expense deleted successfully")
		},
	}

	cmd.Flags().UintVarP(&id, "id", "i", 0, "the expense id that will be deleted")

	return cmd
}

func deleteExpense(id uint, p persistence.ExpensePersister) error {
	if id < 1 {
		return ErrInvalidId
	}

	expenses, err := p.ReadAll()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(expenses, func(e *types.Expense) bool {
		return e.Id == id
	})

	if idx < 0 {
		return ErrExpenseNotFound
	}

	p.Delete(expenses[idx])

	return nil
}
