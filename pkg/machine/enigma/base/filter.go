package base

type Filter struct {
}

func NewFilter() *Filter {
	return new(Filter)
}

func (f *Filter) Filter(c byte) byte {
	if (c >= UpperA) && (c <= UpperZ) {
		return c
	}

	var diff byte = LowerA - UpperA

	if (c >= LowerA) && (c <= LowerZ) {
		return c - diff
	}

	return 0
}
