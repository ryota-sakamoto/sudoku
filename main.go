package main

import (
	"fmt"

	"github.com/ryota-sakamoto/sudoku/sudoku"
)

// const text = `_732__1__
// ____9____
// __615__2_
// ___4___87
// _____3__2
// __49__5__
// __1_7_94_
// _6_______
// _8_______`

const text = `123456789
456789123
789123456
234567891
567891234
891234567
345678912
678912345
91234567_`

func main() {
	t, err := sudoku.Parse(text)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Render())

	if err := t.Solve(); err != nil {
		panic(err)
	}

	fmt.Println(t.Render())
}
