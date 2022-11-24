package context

import "testing"

func TestUseValue(t *testing.T) {
	UseValue()
}

func TestUseCtxValue(t *testing.T) {
	UseCtxValue()
}

func TestUseMap(t *testing.T) {
	UseMap()
}

func TestUseCloseChan(t *testing.T) {
	UseCloseChan()
}

func TestUseParentNil(t *testing.T) {
	UseParentNil()
}

func TestUseParentNotDone(t *testing.T) {
	UseParentDone()
}

func TestUseCtxWrap(t *testing.T) {
	UseCtxWrap()
}
