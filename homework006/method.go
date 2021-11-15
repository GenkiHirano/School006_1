package homework006

import "fmt"

// Triangle 構造体を定義してください。
type Triangle struct {
	Bottom, Height int
}

func (t Triangle) Test_3() int {
	fmt.Println("homework006_3")
	
	return t.Bottom * t.Height / 2
}
