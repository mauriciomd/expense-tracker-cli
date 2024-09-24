package cmd

import (
	"github.com/mauriciomd/expense-tracker/persistence"
	"github.com/spf13/cobra"
)

type rootCommand struct {
	cmd         *cobra.Command
	persistence persistence.ExpensePersister
}

func New(p persistence.ExpensePersister) *rootCommand {

	root := &rootCommand{
		cmd: &cobra.Command{
			Use:   "expense-tracker",
			Short: "A simple expense tracker.",
			Long:  "Expense tracker is a simple cli created to manage your finances.",
		},
		persistence: p,
	}

	root.registerSubCommands()

	return root
}

func (r *rootCommand) registerSubCommands() {
	r.cmd.AddCommand(newAdd(r.persistence))
	r.cmd.AddCommand(newUpdate(r.persistence))
	r.cmd.AddCommand(newDelete(r.persistence))
	r.cmd.AddCommand(newSummary(r.persistence))
	r.cmd.AddCommand(newList(r.persistence))
}

func (r *rootCommand) Execute() error {
	if err := r.cmd.Execute(); err != nil {
		return err
	}

	return nil
}
