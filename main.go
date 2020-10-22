package main

import (
	"fmt"
)

const text = `_732__1__
____9____
__615__2_
___4___87
_____3__2
__49__5__
__1_7_94_
_6_______
_8_______`

func main() {
	t, err := Parse(text)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Render())

	if err := t.Solve(); err != nil {
		panic(err)
	}

	fmt.Println(t.Render())
}
