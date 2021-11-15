package homework006

import "fmt"

type Calculation interface {
	total() int
	average() int
}

type Subject struct {
	Math, English int
}

func (s Subject) total() int {
	return s.Math + s.English
}

func (s Subject) average() int {
	return (s.Math + s.English) / 2
}

func Test_4(c Calculation) {
	fmt.Println("homework006_4_Total")
	fmt.Println(c.Total())
	fmt.Println("homework006_4_Average")
	fmt.Println(c.Average())

}
