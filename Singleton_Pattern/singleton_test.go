package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	s := New()
	s["name"] = "NlogN"

	s1 := New()
	s1["addr"] = "china"

	if (s["name"] != "NlogN") || (s["addr"] != "china") {
		t.Fatal("test faild")
	}
}
