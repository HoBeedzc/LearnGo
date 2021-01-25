package loop

import "testing"

func WhileLoop(t *testing.T) {
	var n int = 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
