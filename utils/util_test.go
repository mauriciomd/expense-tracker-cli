package utils

import (
	"testing"
	"time"

	"github.com/mauriciomd/expense-tracker/types"
)

func TestFilter(t *testing.T) {
	id := uint(1)
	collection := []*types.Expense{
		{Id: id, Description: "another valid description", Amount: 20, Date: time.Now()},
		{Id: id + 1, Description: "another valid description", Amount: 20, Date: time.Now()},
	}

	got := Filter(collection, func(e *types.Expense) bool {
		return e.Id != id
	})

	if len(got) != 1 {
		t.Errorf("got %d but want 0", len(got))
	}
}
