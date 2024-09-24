package cmd

import (
	"errors"
	"testing"
	"time"

	"github.com/mauriciomd/expense-tracker/types"
)

func TestAddExpense(t *testing.T) {
	t.Run("Invalid description", func(t *testing.T) {
		description := ""
		amount := 10.3
		category := "valid category"

		_, got := addExpense(description, category, amount, &MockPersistence{})
		assertError(t, got, ErrInvalidDescription)
	})

	t.Run("Invalid amount", func(t *testing.T) {
		description := "breakfast"
		amount := -10.0
		category := "valid category"

		_, got := addExpense(description, category, amount, &MockPersistence{})
		assertError(t, got, ErrInvalidAmount)
	})

	t.Run("Default category", func(t *testing.T) {
		cases := []struct {
			description string
			amount      float64
			category    string
		}{
			{description: "breakfast", amount: 10.0, category: ""},
			{description: "breakfast", amount: 10.0, category: " "},
		}

		for _, test := range cases {
			got, err := addExpense(test.description, test.category, test.amount, &MockPersistence{})

			assertErrorIsNil(t, err)
			assertAny(t, got.Category, defaultCategory)
		}
	})

	t.Run("Specific category", func(t *testing.T) {
		cases := []struct {
			description string
			amount      float64
			category    string
		}{
			{description: "breakfast", amount: 10.0, category: "food"},
		}

		for _, test := range cases {
			got, err := addExpense(test.description, test.category, test.amount, &MockPersistence{})

			assertErrorIsNil(t, err)
			assertAny(t, got.Category, test.category)
		}
	})

	t.Run("Add new expense", func(t *testing.T) {
		want := types.Expense{
			Id:          1,
			Description: "Breakfast",
			Category:    "food",
			Amount:      10,
			Date:        time.Now().Local(),
		}

		mp := &MockPersistence{}

		got, err := addExpense(
			want.Description,
			want.Category,
			want.Amount,
			mp)

		assertErrorIsNil(t, err)
		assertId(t, got.Id, want.Id)
		assertAny(t, got.Category, want.Category)
		assertAny(t, got.Description, want.Description)
		assertAny(t, got.Amount, want.Amount)
		assertDate(t, got.Date, want.Date)

		if len(mp.data) == 0 {
			t.Fatal("data hasn't been persisted")
		}
	})
}

func TestGetNextId(t *testing.T) {
	t.Run("First expense", func(t *testing.T) {
		want := uint(1)
		got, err := getNextId(&MockPersistence{})
		assertErrorIsNil(t, err)
		assertId(t, got, want)
	})

	t.Run("Next sequential id", func(t *testing.T) {
		mock := &MockPersistence{
			data: []*types.Expense{{Id: 10}},
		}
		want := 11
		got, err := getNextId(mock)
		assertErrorIsNil(t, err)
		assertId(t, got, uint(want))
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("got nil but want %q", want.Error())
	}

	if !errors.Is(got, want) {
		t.Fatalf("got '%q' but want '%q'", got, ErrInvalidDescription)
	}
}

func assertErrorIsNil(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("got error '%q' but want nil", got.Error())
	}
}

func assertId(t *testing.T, got, want uint) {
	t.Helper()
	if got != want {
		t.Fatalf("got '%d' but want '%d'", got, want)
	}
}

func assertAny(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Fatalf("got '%v' but want '%v'", got, want)
	}
}

func assertDate(t *testing.T, got, want time.Time) {
	t.Helper()
	format := "2006-01-02"
	formattedGot := got.Format(format)
	formattedWant := want.Format(format)

	if formattedGot != formattedWant {
		t.Fatalf("got '%v' but want '%v'", got, want)
	}
}
