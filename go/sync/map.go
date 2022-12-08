package sync

import (
	"fmt"
	"sync"
)

func Map() {
	m := sync.Map{}
	m.Store("a", "v")
	fmt.Println(m.Load("a"))
}
