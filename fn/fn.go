package fn

import "fmt"

//AcceptAnonymousFunc - на вход принимает функцию вида A func(int,int)int, а внутри оборачивает и вызывает её при выходе
func AcceptAnonymousFunc(A func(x int, y int) int) {
	fmt.Println("**************")
	fmt.Println("Start function")
	x := 2
	y := 8
	defer fmt.Println("Result:", A(x, y))
	fmt.Println("Something is going on")
}
