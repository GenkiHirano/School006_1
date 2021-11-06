package homework006

import "fmt"

type Triangle struct {
	Bottom, Height int
}

func (r *Triangle) Triangle() int {
	return r.Bottom * r.Height / 2
}

func main() {
	r := Triangle{Bottom: 25, Height: 10}

	fmt.Println("Triangle: ", r.Triangle())
}
