package slice

import (
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInfo(t *testing.T) {
	cases := []struct {
		in   int
		len  int
		cap  int
		name string
	}{
		{
			in:   0,
			len:  0,
			cap:  0,
			name: "0",
		},
		{
			in:   1,
			len:  1,
			cap:  2,
			name: "1-小于1024则是2倍",
		},
		{
			in:   2,
			len:  2,
			cap:  2,
			name: "2-小于1024则是2倍，但cap够用不扩容",
		},
		{
			in:   3,
			len:  3,
			cap:  4,
			name: "3-小于1024则是2倍，扩容两次，4",
		},
		{
			in:   1000,
			len:  1000,
			cap:  1024,
			name: "1000",
		},
	}
	for _, v := range cases {
		convey.Convey(v.name, t, func() {
			l, c := Info(v.in)
			convey.So(l, assertions.ShouldEqual, v.len)
			convey.So(c, assertions.ShouldEqual, v.cap)
		})
	}
}

func BenchmarkGeneCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneCap(10)
	}
}

func BenchmarkGeneNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneNoCap(10)
	}
}
