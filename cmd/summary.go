package cmd

import (
	"fmt"
	"time"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
	"github.com/spf13/cobra"
)

var summaryFmt = "Total expenses: $%.2f\n"
var summarySpecificMonthFmt = "Total expenses for %q: $%.2f\n"

func newSummary(p persistence.ExpensePersister) *cobra.Command {
	var month int

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Show a summary of all expenses",
		Run: func(cmd *cobra.Command, args []string) {
			amount, err := summary(month, p)
			if err != nil {
				fmt.Printf("error: %q", err.Error())
				return
			}

			fmt.Println(formatMessage(amount, month))
		},
	}

	cmd.Flags().IntVarP(&month, "month", "m", 0, "Show a summary of expenses (optionally, for a specific month of current year)")

	return cmd
}

func summary(month int, p persistence.ExpensePersister) (float64, error) {
	if month < 0 || month > 12 {
		return 0, ErrInvalidMonth
	}

	expenses, err := p.ReadAll()
	if err != nil {
		return 0, err
	}

	expenses = filterByCurrentYear(expenses)
	if month > 0 {
		expenses = filterByMonth(expenses, month)
	}

	sum := 0.0
	for _, e := range expenses {
		sum += e.Amount
	}
	return sum, nil
}

func filterByCurrentYear(data []*types.Expense) []*types.Expense {
	var filteredExpenses []*types.Expense
	currentYear := time.Now().Year()

	for _, e := range data {
		if e.Date.Year() == currentYear {
			filteredExpenses = append(filteredExpenses, e)
		}
	}

	return filteredExpenses
}

func filterByMonth(data []*types.Expense, m int) []*types.Expense {
	var filteredExpenses []*types.Expense
	month := time.Month(m)

	for _, e := range data {
		if e.Date.Month() == month {
			filteredExpenses = append(filteredExpenses, e)
		}
	}

	return filteredExpenses
}

func formatMessage(amount float64, month int) string {
	if month == 0 {
		return fmt.Sprintf(summaryFmt, amount)
	}

	return fmt.Sprintf(summarySpecificMonthFmt, time.Month(month), amount)
}
