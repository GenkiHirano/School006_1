package homework006

import "fmt"

type User struct {
	name  string
	score int
}

func Test_2() {
	u := new(User)
	u.name = "佐藤"
	u.score = 100

	fmt.Println(u.name)
	fmt.Println(u.score)
}
