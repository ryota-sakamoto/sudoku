package sudoku

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

// Table is sudoku's table
type Table struct {
	Panels [9][9]map[int]struct{}
}

func Parse(s string) (*Table, error) {
	rows := strings.Split(s, "\n")
	if len(rows) != 9 {
		return nil, fmt.Errorf("invalid rows count: want 9, but %d", len(rows))
	}

	panels := [9][9]map[int]struct{}{}

	for i, v := range rows {
		if len(v) != 9 {
			return nil, fmt.Errorf("invalid line length(%d): want 9, but %d", i+1, len(v))
		}

		for j, p := range v {
			if (p < 49 || 57 < p) && p != 95 {
				return nil, fmt.Errorf("invalid char: want ([1-9]|_), but %s", string(p))
			}

			if panels[i][j] == nil {
				panels[i][j] = map[int]struct{}{}
			}

			if p != 95 {
				panels[i][j][int(p)-48] = struct{}{}
			}
		}
	}

	for i := range panels {
		for j := range panels[i] {
			if len(panels[i][j]) == 1 {
				continue
			}

			p := map[int]struct{}{}
			for k := 0; k < 9; k++ {
				if len(panels[i][k]) != 1 {
					continue
				}

				for key := range panels[i][k] {
					p[key] = struct{}{}
				}
			}

			for k := 0; k < 9; k++ {
				if len(panels[k][j]) != 1 {
					continue
				}

				for key := range panels[k][j] {
					p[key] = struct{}{}
				}
			}

			_i := i / 3 * 3
			_j := j / 3 * 3
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if len(panels[_i+x][_j+y]) != 1 {
						continue
					}

					for key := range panels[_i+x][_j+y] {
						p[key] = struct{}{}
					}
				}
			}

			for k := 1; k <= 9; k++ {
				if _, ok := p[k]; !ok {
					panels[i][j][k] = struct{}{}
				}
			}
		}
	}

	return &Table{
		Panels: panels,
	}, nil
}

func (t *Table) Solve() error {
	return t.dfs(0, 0)
}

// TODO: improvement
func (t *Table) dfs(i, j int) error {
	if i == 9 && j == 0 {
		return t.Check()
	}

	t.Debug()

	raw := map[int]struct{}{}
	for key := range t.Panels[i][j] {
		raw[key] = struct{}{}
	}

	if len(raw) == 0 {
		return fmt.Errorf("candidate count is zero")
	}

	for key := range raw {
		t.Panels[i][j] = map[int]struct{}{
			key: {},
		}

		nexti := i
		nextj := j

		if j == 8 {
			nexti++
			nextj = 0
		} else {
			nextj++
		}

		before := t.Removal(i, j, key)
		fmt.Println("input:", i, j, key)
		fmt.Println("before:", before)
		fmt.Println()

		if err := t.dfs(nexti, nextj); err == nil {
			return nil
		}

		for _, v := range before {
			t.Panels[v.x][v.y][key] = struct{}{}
		}
	}

	return fmt.Errorf("no")
}

type XY struct {
	x int
	y int
}

func (t *Table) Removal(i, j, num int) []XY {
	result := []XY{}

	for k := 0; k < 9; k++ {
		if j == k {
			continue
		}

		if _, ok := t.Panels[i][k][num]; ok {
			result = append(result, XY{
				x: i,
				y: k,
			})

			delete(t.Panels[i][k], num)
		}
	}

	for k := 0; k < 9; k++ {
		if i == k {
			continue
		}

		if _, ok := t.Panels[k][j][num]; ok {
			result = append(result, XY{
				x: k,
				y: j,
			})

			delete(t.Panels[k][j], num)
		}
	}

	_i := i / 3 * 3
	_j := j / 3 * 3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if _i+x == i && _j+y == j {
				continue
			}

			if _, ok := t.Panels[_i+x][_j+y][num]; ok {
				result = append(result, XY{
					x: _i + x,
					y: _j + y,
				})

				delete(t.Panels[_i+x][_j+y], num)
			}
		}
	}

	return result
}

func (t *Table) Check() error {
	return nil
}

func (t *Table) Render() string {
	writer := table.NewWriter()

	l := []string{}
	for i := range t.Panels {
		s := ""
		for j := range t.Panels[i] {
			if len(t.Panels[i][j]) == 1 {
				for key := range t.Panels[i][j] {
					s += strconv.Itoa(key)
				}
			} else {
				s += "_"
			}
		}

		l = append(l, s)
		if len(l) == 3 {
			r := table.Row{}
			for k := 0; k < 9; k += 3 {
				a := string(l[0][k]) + string(l[0][k+1]) + string(l[0][k+2])
				b := string(l[1][k]) + string(l[1][k+1]) + string(l[1][k+2])
				c := string(l[2][k]) + string(l[2][k+1]) + string(l[2][k+2])

				r = append(r, a+"\n"+b+"\n"+c)
			}
			writer.AppendRow(r)
			l = []string{}
		}
	}

	writer.Style().Options.SeparateRows = true

	return writer.Render()
}

func (t *Table) Debug() {
	s := ""
	for i := range t.Panels {
		for j := range t.Panels[i] {
			if len(t.Panels[i][j]) == 1 {
				for key := range t.Panels[i][j] {
					s += strconv.Itoa(key)
				}
			} else {
				s += "_"
			}
		}
	}
	fmt.Println(s)
}
