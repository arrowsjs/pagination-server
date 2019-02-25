package db

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/efritz/nacelle"
)

type Initializer struct {
	Logger    nacelle.Logger           `service:"logger"`
	Container nacelle.ServiceContainer `service:"container"`
}

const ServiceName = "db"

func NewInitializer() *Initializer {
	return &Initializer{}
}

func (i *Initializer) Init(config nacelle.Config) error {
	records, err := readRecords("data/db.csv")
	if err != nil {
		return err
	}

	items := []Item{}
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		pricePerUnit, _ := strconv.ParseFloat(record[2], 64)
		margin, _ := strconv.ParseFloat(record[6], 64)

		items = append(items, Item{
			ID:            id,
			Name:          record[1],
			Category:      record[3],
			SubCategory:   record[4],
			ContainerType: record[5],
			PricePerUnit:  pricePerUnit,
			Margin:        margin,
		})
	}

	i.Logger.Info("Loaded %d records from disk.", len(items))
	return i.Container.Set(ServiceName, NewDB(items))
}

//
// Helpers

func readRecords(path string) ([][]string, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(r).ReadAll()
}
