package atomic

import "testing"

func TestAdd(t *testing.T) {
	t.Log(Add(1000))
}

func TestCompareAndSwap(t *testing.T) {
	if 1 != CompareAndSwap(1000) {
		t.Fatal("err")
	}
}
