package slice

func Info(num int) (sLen, sCap int) {
	s := make([]int, 0)
	for i := 0; i < num; i++ {
		s = append(s, i)
	}
	return len(s), cap(s)
}

func GeneCap(num int) {
	s := make([]int, 0, num)
	for i := 0; i < num; i++ {
		s = append(s, i)
	}
}

func GeneNoCap(num int) {
	s := make([]int, 0)
	for i := 0; i < num; i++ {
		s = append(s, i)
	}
}
