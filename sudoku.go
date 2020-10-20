package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

// Table is sudoku's table
type Table struct {
	_panels [9][9][]int
}

func Parse(s string) (*Table, error) {
	rows := strings.Split(s, "\n")
	if len(rows) != 9 {
		return nil, fmt.Errorf("invalid rows count: want 9, but %d", len(rows))
	}

	panels := [][]map[int]int{
		{
			{},
			{},
			{},
		},
		{
			{},
			{},
			{},
		},
		{
			{},
			{},
			{},
		},
	}

	for i, v := range rows {
		if len(v) != 9 {
			return nil, fmt.Errorf("invalid line length(%d): want 9, but %d", i+1, len(v))
		}

		for j, p := range v {
			if (p < 49 || 57 < p) && p != 95 {
				return nil, fmt.Errorf("invalid char: want ([1-9]|_), but %s", string(p))
			}

			if p == 95 {
				panels[i/3][j/3][i%3*3+j%3] = 0
			} else {
				panels[i/3][j/3][i%3*3+j%3] = int(p) - 48
			}
		}
	}

	var _panels [9][9][]int
	for i := range panels {
		for j := range panels[i] {
			for k, v := range panels[i][j] {
				if v != 0 {
					_panels[i*3+j][k] = append(_panels[i*3+j][k], v)
				} else {
					// _panels[i*3+j][k] = append(_panels[i*3+j][k])
				}
			}
		}
	}

	return &Table{
		_panels: _panels,
	}, nil
}

func (t *Table) Solve() error {
	return nil
}

func (t *Table) dfs() {

}

func (t *Table) Render() string {
	writer := table.NewWriter()

	r := table.Row{}
	for _, v := range t._panels {
		s := ""
		for _, p := range v {
			if len(p) == 1 {
				s += strconv.Itoa(p[0])
			} else {
				s += "_"
			}

			if len(s) == 3 || len(s) == 7 {
				s += "\n"
			}

			if len(s) == 11 {
				r = append(r, s)
			}
		}

		if len(r) == 3 {
			writer.AppendRow(r)
			r = table.Row{}
		}
	}

	writer.Style().Options.SeparateRows = true

	return writer.Render()
}
