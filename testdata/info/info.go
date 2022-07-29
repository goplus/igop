package info

type Info interface {
	Mode() int
	SetMode(n int)
	Count(a ...int) int
}

type MyInfo struct {
	mode int
}

func (i MyInfo) Mode() int {
	return i.mode
}

func (i *MyInfo) SetMode(mode int) {
	i.mode = mode
}

func (i *MyInfo) Count(a ...int) int {
	return len(a)
}

func NewInfo(mode int) *MyInfo {
	return &MyInfo{mode}
}
