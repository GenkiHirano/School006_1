package homework006

import "fmt"

// User 構造体を定義してください。
//
type User struct {
	name  string
	score int
}

func Test_2() {
	fmt.Println("homework006_2")
	// 第6回スクール 宿題2 構造体を実装してください。
	fmt.Println(User{name: "佐藤", score: 100})
}
