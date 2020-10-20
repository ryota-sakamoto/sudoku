package main

import (
	"fmt"
)

const text = `___9__6__
4_____8__
2_76__5_4
1_87__3__
_____3___
_______4_
_9___7__1
__5__2___
_3_5_5___`

func main() {
	t, err := Parse(text)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Render())
}
