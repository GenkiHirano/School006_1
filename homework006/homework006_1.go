package homework006

import "fmt"

func Test_1() {

	s := make(map[string]int)

	s["佐藤"] = 85
	s["鈴木"] = 65

	fmt.Println(s)

	delete(s, "佐藤")
	fmt.Println(s)
}
