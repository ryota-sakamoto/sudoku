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

var lrtable = map[int]map[rune][]int{
	0: {
		'r': []int{1, 2},
		'l': []int{3, 6},
	},
	1: {
		'r': []int{0, 2},
		'l': []int{4, 7},
	},
	2: {
		'r': []int{0, 1},
		'l': []int{5, 8},
	},
	3: {
		'r': []int{4, 5},
		'l': []int{0, 6},
	},
	4: {
		'r': []int{3, 5},
		'l': []int{1, 7},
	},
	5: {
		'r': []int{3, 4},
		'l': []int{2, 8},
	},
	6: {
		'r': []int{7, 8},
		'l': []int{0, 3},
	},
	7: {
		'r': []int{6, 8},
		'l': []int{1, 4},
	},
	8: {
		'r': []int{6, 7},
		'l': []int{2, 5},
	},
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
			p := map[int]struct{}{}
			for k, v := range panels[i][j] {
				if v != 0 {
					_panels[i*3+j][k] = append(_panels[i*3+j][k], v)
					p[v] = struct{}{}
				}
			}

			ac := []int{}
			for c := 1; c <= 9; c++ {
				if _, ok := p[c]; !ok {
					ac = append(ac, c)
				}
			}

			for k, v := range panels[i][j] {
				if v == 0 {
					_panels[i*3+j][k] = make([]int, len(ac))
					copy(_panels[i*3+j][k], ac)
				}
			}
		}
	}

	return &Table{
		_panels: _panels,
	}, nil
}

func (t *Table) RemovalInPanel(i int) {
	p := map[int]struct{}{}
	for j := range t._panels[i] {
		if len(t._panels[i][j]) == 1 {
			p[t._panels[i][j][0]] = struct{}{}
		}
	}

	ac := []int{}
	for c := 1; c <= 9; c++ {
		if _, ok := p[c]; !ok {
			ac = append(ac, c)
		}
	}

	for j := range t._panels[i] {
		if len(t._panels[i][j]) != 1 {
			t._panels[i][j] = make([]int, len(ac))
			copy(t._panels[i][j], ac)
		}
	}
}

func (t *Table) Removal() {
	for i := range t._panels {
		for j := range t._panels[i] {
			if len(t._panels[i][j]) == 1 {
				continue
			}

			for _, p := range lrtable[i]['r'] {
				for k := 0; k < 3; k++ {
					t2 := t._panels[p][j/3*3+k]
					if len(t2) == 1 {
						p := []int{}
						for _, c := range t._panels[i][j] {
							if t2[0] != c {
								p = append(p, c)
							}
						}
						t._panels[i][j] = p
					}
				}
			}

			for _, p := range lrtable[i]['l'] {
				for k := 0; k < 3; k++ {
					t2 := t._panels[p][j%3+3*k]
					if len(t2) == 1 {
						p := []int{}
						for _, c := range t._panels[i][j] {
							if t2[0] != c {
								p = append(p, c)
							}
						}
						t._panels[i][j] = p
					}
				}
			}
		}
	}
}

func (t *Table) Solve() error {
	t.Removal()
	return t.dfs(0, 0)
}

// TODO: improvement
func (t *Table) dfs(i, j int) error {
	if i == 9 && j == 0 {
		return t.Check()
	}

	t.RemovalInPanel(i)
	t.Removal()

	raw := make([]int, len(t._panels[i][j]))
	copy(raw, t._panels[i][j])

	for _, can := range raw {
		t._panels[i][j] = []int{can}

		nexti := i
		nextj := j

		if j == 8 {
			nexti++
			nextj = 0
		} else {
			nextj++
		}

		if err := t.dfs(nexti, nextj); err == nil {
			return nil
		}
	}

	t._panels[i][j] = raw

	return fmt.Errorf("no")
}

func (t *Table) Check() error {
	for i := range t._panels {
		al := map[int]struct{}{}
		for j := 0; j < 9; j++ {
			p := t._panels[i][j]
			if len(p) != 1 {
				return fmt.Errorf("invalid count")
			}

			point := p[0]
			if _, ok := al[point]; !ok {
				al[point] = struct{}{}
			} else {
				return fmt.Errorf("already in panel")
			}
		}

		for j := range t._panels[i] {
			for _, p := range lrtable[i]['r'] {
				for k := 0; k < 3; k++ {
					c := t._panels[p][j/3*3+k]
					if t._panels[i][j][0] == c[0] {
						return fmt.Errorf("already in row")
					}
				}
			}
		}

		for j := range t._panels[i] {
			for _, p := range lrtable[i]['l'] {
				for k := 0; k < 3; k++ {
					c := t._panels[p][j%3+3*k]
					if t._panels[i][j][0] == c[0] {
						return fmt.Errorf("already in line")
					}
				}
			}
		}
	}

	return nil
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

func (t *Table) Yoko() string {
	s := ""
	for i := range t._panels {
		for j := range t._panels[i] {
			s += strconv.Itoa(t._panels[i][j][0])
		}
	}

	return s
}
