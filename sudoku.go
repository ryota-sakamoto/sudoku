package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

// Panel is 9x9 squares
// 1 2 3
// 4 5 6
// 7 8 9
type Panel map[int]int

// Table is sudoku's table
// panels[0][0] panels[0][1] panels[0][2]
// panels[1][0] panels[1][1] panels[1][2]
// panels[2][0] panels[2][1] panels[1][2]
type Table struct {
	panels [][]Panel
}

func Parse(s string) (*Table, error) {
	rows := strings.Split(s, "\n")
	if len(rows) != 9 {
		return nil, fmt.Errorf("invalid rows count: want 9, but %d", len(rows))
	}

	panels := [][]Panel{
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

	return &Table{
		panels: panels,
	}, nil
}

func (t *Table) Solve() error {
	return nil
}

func (t *Table) dfs() {

}

func (t *Table) Render() string {
	writer := table.NewWriter()

	for _, v := range t.panels {
		r := table.Row{}
		for _, p := range v {
			s := ""
			for i := 0; i < 9; i++ {
				if p[i] == 0 {
					s += "_"
				} else {
					s += strconv.Itoa(p[i])
				}

				if len(s) == 3 || len(s) == 7 {
					s += "\n"
				}
			}
			r = append(r, s)
		}

		if len(r) == 3 {
			writer.AppendRow(r)
		}
	}

	writer.Style().Options.SeparateRows = true

	return writer.Render()
}
