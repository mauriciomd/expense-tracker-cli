package cmd

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/mauriciomd/expense-tracker/types"
)

func TestSummary(t *testing.T) {
	t.Run("invalid month", func(t *testing.T) {
		cases := []struct {
			month int
			want  error
		}{
			{month: -1, want: ErrInvalidMonth},
			{month: 13, want: ErrInvalidMonth},
		}

		for _, test := range cases {
			_, got := summary(test.month, &persistence.MockPersistence{})

			if got != test.want {
				t.Fatalf("got %q want %q", got, test.want)
			}
		}
	})

	t.Run("all expenses of current year", func(t *testing.T) {
		lastYear := time.Date(time.Now().Year()-1, 1, 1, 0, 0, 0, 0, time.Local)

		mock := &persistence.MockPersistence{
			Data: []*types.Expense{
				{Id: 1, Description: "breakfast", Amount: 10, Date: time.Now()},
				{Id: 2, Description: "lunch", Amount: 30, Date: lastYear},
				{Id: 3, Description: "dinner", Amount: 20, Date: time.Now()},
			},
		}

		month := 0
		want := 30.0
		got, _ := summary(month, mock)

		if got != want {
			t.Fatalf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("with a specific month", func(t *testing.T) {
		month := 2
		now := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Local)
		lastMonth := time.Date(time.Now().Year(), time.Month(month-1), 1, 0, 0, 0, 0, time.Local)

		mock := &persistence.MockPersistence{
			Data: []*types.Expense{
				{Id: 1, Description: "breakfast", Amount: 20, Date: now},
				{Id: 3, Description: "dinner", Amount: 45, Date: now},
				{Id: 2, Description: "lunch", Amount: 50, Date: lastMonth},
			},
		}

		want := 65.0
		got, _ := summary(month, mock)

		if got != want {
			t.Fatalf("got %.2f want %.2f", got, want)
		}
	})
}

func TestFormatMessage(t *testing.T) {
	t.Run("without a specific month", func(t *testing.T) {
		amount := 10.0
		month := 0
		want := fmt.Sprintf(summaryFmt, amount)
		got := formatMessage(amount, month)

		if got != want {
			t.Fatalf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("without a specific month", func(t *testing.T) {
		amount := 10.0
		month := 1
		want := fmt.Sprintf(summarySpecificMonthFmt, time.Month(month), amount)
		got := formatMessage(amount, month)

		if got != want {
			t.Fatalf("got '%s' want '%s'", got, want)
		}
	})
}

func TestFilterByCurrentYear(t *testing.T) {
	lastYear := time.Date(time.Now().Year()-1, 1, 1, 0, 0, 0, 0, time.Local)
	data := []*types.Expense{
		{Id: 1, Description: "breakfast", Amount: 10, Date: time.Now()},
		{Id: 2, Description: "lunch", Amount: 30, Date: lastYear},
		{Id: 3, Description: "dinner", Amount: 20, Date: time.Now()},
	}

	want := []*types.Expense{
		{Id: 1, Description: "breakfast", Amount: 10, Date: time.Now()},
		{Id: 3, Description: "dinner", Amount: 20, Date: time.Now()},
	}

	got := filterByCurrentYear(data)
	cmp := slices.CompareFunc(got, want, func(a, b *types.Expense) int {
		if a.Id != b.Id {
			return 1
		}

		return 0
	})

	if cmp != 0 {
		t.Fatalf("got '%v' want '%v'", got, want)
	}
}

func TestFilterMonth(t *testing.T) {
	month := 2
	currentDate := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastMonth := time.Date(time.Now().Year(), time.Month(month-1), 1, 0, 0, 0, 0, time.Local)

	data := []*types.Expense{
		{Id: 1, Description: "breakfast", Amount: 10, Date: currentDate},
		{Id: 2, Description: "lunch", Amount: 30, Date: lastMonth},
	}

	want := []*types.Expense{
		{Id: 1, Description: "breakfast", Amount: 10, Date: currentDate},
	}

	got := filterByMonth(data, month)
	cmp := slices.CompareFunc(got, want, func(a, b *types.Expense) int {
		if a.Id != b.Id {
			return 1
		}

		return 0
	})

	if cmp != 0 {
		t.Fatalf("got '%v' want '%v'", got, want)
	}
}
