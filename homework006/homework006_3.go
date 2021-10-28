package homework006

import "fmt"

type Student struct {
	name          string
	math, english int
}

func (s Student) Test_3() {
	fmt.Println(s.name, (s.math+s.english)/2)
}
