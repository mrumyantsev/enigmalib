package machine

type Filter struct {
}

func NewFilter() *Filter {
	return &Filter{}
}

func EmptyFilter() *Reflector {
	return &Reflector{}
}

func (f *Filter) Filter(c byte) byte {
	if (c >= 'A') && (c <= 'Z') {
		return c
	}

	var diff byte = 'a' - 'A'

	if (c >= 'a') && (c <= 'z') {
		return c - diff
	}

	return 0
}
