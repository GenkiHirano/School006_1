package main

import (
	"fmt"

	"github.com/GenkiHirano/homework006/homework006"
)

func main() {
	homework006.Test_1()
	homework006.Test_2()
	t := &homework006.Triangle{
		Bottom: 25,
		Height: 10,
	}

	fmt.Println(t.Test_3())

	s := &homework006.Subject{
		Math:    90,
		English: 70,
	}
	homework006.Test_4(s)
}
