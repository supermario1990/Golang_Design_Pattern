package pool

import (
	"fmt"
	"testing"
)

var (
	count int
)

func TestPool(t *testing.T) {
	p := New(2)
	select {
	case obj := <-(*p):
		obj.Desc = "hello"
		fmt.Println(obj)
		(*p) <- obj
		fmt.Println("p:", p)
	default:
		return

	}
}
