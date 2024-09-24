package persistence

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mauriciomd/expense-tracker/types"
	"github.com/mauriciomd/expense-tracker/utils"
)

type FilePersistence struct {
	filename string
}

func NewFilePersistence(filename string) (*FilePersistence, error) {
	p := &FilePersistence{
		filename,
	}

	if err := p.createFileIfNotExists(); err != nil {
		return nil, err
	}

	return p, nil
}

func (fp *FilePersistence) Add(e *types.Expense) error {
	file, err := os.OpenFile(fp.filename, os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	id := strconv.Itoa(int(e.Id))
	amount := strconv.FormatFloat(e.Amount, 'f', -1, 64)
	year, month, day := e.Date.Date()
	date := fmt.Sprintf("%d-%d-%d", year, month, day)
	writer := csv.NewWriter(file)

	defer writer.Flush()

	err = writer.Write([]string{id, e.Description, e.Category, amount, date})
	if err != nil {
		return err
	}

	return nil
}

func (fp *FilePersistence) Delete(e *types.Expense) error {
	if e == nil {
		return errors.New("invalid argument")
	}

	expenses, err := fp.ReadAll()
	if err != nil {
		return err
	}

	expenses = utils.Filter(expenses, func(exp *types.Expense) bool {
		return e.Id != exp.Id
	})

	if err := os.Truncate(fp.filename, 0); err != nil {
		return nil
	}

	for _, expense := range expenses {
		if err := fp.Add(expense); err != nil {
			continue
		}
	}

	return nil
}

func (fp *FilePersistence) ReadAll() ([]*types.Expense, error) {
	file, err := os.Open(fp.filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	output := make([]*types.Expense, len(records))
	for i, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		amount, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		date, err := time.Parse("2006-1-02", record[4])
		if err != nil {
			continue
		}

		expense := &types.Expense{
			Id:          uint(id),
			Description: record[1],
			Category:    record[2],
			Amount:      amount,
			Date:        date,
		}

		output[i] = expense
	}

	return output, nil
}

func (fp *FilePersistence) Update(e *types.Expense) error {
	if err := fp.Delete(e); err != nil {
		return err
	}

	if err := fp.Add(e); err != nil {
		return err
	}

	return nil
}

func (fp *FilePersistence) createFileIfNotExists() error {
	file, err := os.Open(fp.filename)
	if err == nil || !errors.Is(err, os.ErrNotExist) {
		return nil
	}

	defer file.Close()

	file, err = os.Create(fp.filename)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}
