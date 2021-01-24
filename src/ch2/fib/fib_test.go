package fib

import "testing"

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		fmt.Println(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Plintln(b)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)
}
