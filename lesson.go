package main

import "fmt"

// 関数外でも宣言できる
// 型を明示的に指定したいときはvarを使う
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	// ↓の方法では、関数内でしか宣言できない
	xi := 1
	xi = 2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false

	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func main() {
	foo()
	fmt.Println(i, f64, s, t, f)
}
