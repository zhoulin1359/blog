package atomic

import "testing"

func TestAdd(t *testing.T) {
	t.Log(Add(1000))
}

//https://amazingao.com/posts/2020/11/go-src/sync/atomic/
//https://golang.design/under-the-hood/zh-cn/part1basic/ch01basic/asm/
func TestCompareAndSwap(t *testing.T) {
	if 1 != CompareAndSwap(1000) {
		t.Fatal("err")
	}
}
