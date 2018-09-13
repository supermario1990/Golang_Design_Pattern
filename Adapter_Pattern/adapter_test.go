package adapter

import "testing"

var expect = "adaptee method"

func TestAdaper(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect:%s, res:%s", expect, res)
	}
}
