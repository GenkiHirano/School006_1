package homework006

import "fmt"

// 第6回スクール 宿題4 インターフェースを実装してください。
//
type Calculation interface {
	Total() int
	Average() int
}
type Subject struct {
	Math, English int
}

func (s Subject) Total() int {
	return s.Math + s.English
}

func (s Subject) Average() int {
	return (s.Math + s.English) / 2
}

func Test_4(c Calculation) {
	fmt.Println("homework006_4_Total")
	fmt.Println(c.Total())
	fmt.Println("homework006_4_Average")
	fmt.Println(c.Average())
}
