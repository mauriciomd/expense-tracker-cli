package types

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          uint
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

func (e Expense) String() string {
	fmtDate := e.Date.Format("2006-01-02")
	return fmt.Sprintf("{\nid: %d, \ndescription: %s, \namount %.2f, \ndate: %s \n}", e.Id, e.Description, e.Amount, fmtDate)
}
