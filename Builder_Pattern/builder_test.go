package builder

import (
	"testing"
)

func TestBuiler(t *testing.T) {
	bd := &Builder1{}
	director := NewDirector(bd)
	director.Construct()
	res := bd.GetResult()
	if res != "123" {
		t.Fatal("test failed!")
	}
}
