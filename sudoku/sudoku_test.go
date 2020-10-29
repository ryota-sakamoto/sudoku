package sudoku_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ryota-sakamoto/sudoku/sudoku"
	"github.com/ryota-sakamoto/sudoku/testutil"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		result *sudoku.Table
	}{
		{
			s: `123456789
			456789123
			789123456
			234567891
			567891234
			891234567
			345678912
			678912345
			91234567_`,
			result: &sudoku.Table{
				Panels: [9][9]map[int]struct{}{
					testutil.CreateSequence(1, 2, 3, 4, 5, 6, 7, 8, 9),
					testutil.CreateSequence(4, 5, 6, 7, 8, 9, 1, 2, 3),
					testutil.CreateSequence(7, 8, 9, 1, 2, 3, 4, 5, 6),
					testutil.CreateSequence(2, 3, 4, 5, 6, 7, 8, 9, 1),
					testutil.CreateSequence(5, 6, 7, 8, 9, 1, 2, 3, 4),
					testutil.CreateSequence(8, 9, 1, 2, 3, 4, 5, 6, 7),
					testutil.CreateSequence(3, 4, 5, 6, 7, 8, 9, 1, 2),
					testutil.CreateSequence(6, 7, 8, 9, 1, 2, 3, 4, 5),
					testutil.CreateSequence(9, 1, 2, 3, 4, 5, 6, 7, 8),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := sudoku.Parse(strings.Replace(tt.s, "\t", "", -1))
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, tt.result, result)
		})
	}
}
