package homework006

import "fmt"

// 第6回スクール 宿題4 インターフェースを実装してください。

func Test_4(c Calculation) {
	fmt.Println("homework006_4_Total")
	fmt.Println(c.Total())
	fmt.Println("homework006_4_Average")
	fmt.Println(c.Average())
}
