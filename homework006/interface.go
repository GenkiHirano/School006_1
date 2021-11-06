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

func Calculation2(c Calculation) {
	fmt.Println("2教科の合計：", c.total())
	fmt.Println("2教科の平均：", c.average())
}

func main() {
	var a, b int
	fmt.Printf("数学の点数：")
	fmt.Scan(&a)
	fmt.Printf("英語の点数：")
	fmt.Scan(&b)
	s := Subject{Math: a, English: b}

	Calculation2(s)
}
