package context

import (
	"fmt"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(time.Second * 10)
}
