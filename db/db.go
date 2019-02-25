package db

import "strings"

type (
	DB struct {
		items []Item
	}

	Result struct {
		Query      string `json:"query"`
		Results    []Item `json:"results"`
		Prev       int    `json:"prev"`
		Next       int    `json:"next"`
		RangeLeft  int    `json:"rangeLeft"`
		RangeRight int    `json:"rangeRight"`
		Count      int    `json:"count"`
	}
)

func NewDB(items []Item) *DB {
	return &DB{
		items: items,
	}
}

func (db *DB) Search(q string, page, pageSize int) Result {
	results := []Item{}
	for _, record := range db.items {
		if strings.Contains(strings.ToLower(record.Name), strings.ToLower(q)) {
			results = append(results, record)
		}
	}

	m := min(len(results), (page-1)*pageSize)
	n := min(len(results), (page-0)*pageSize)

	left := m + 1
	if len(results) == 0 {
		left = 0
	}

	return Result{
		Query:      q,
		Results:    results[m:n],
		Prev:       page - 1,
		Next:       page + 1,
		RangeLeft:  left,
		RangeRight: n,
		Count:      len(results),
	}
}
