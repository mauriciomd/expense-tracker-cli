package main

import (
	"github.com/mauriciomd/expense-tracker/cmd"
	"github.com/mauriciomd/expense-tracker/persistence"
)

func main() {
	persistence, err := persistence.NewFilePersistence("expenses.csv")
	if err != nil {
		panic(err)
	}

	cli := cmd.New(persistence)
	cli.Execute()
}
