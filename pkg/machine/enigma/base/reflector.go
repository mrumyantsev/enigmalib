package base

type Reflector struct {
	wiring     []byte
	charsCount byte
}

func NewReflector(encoding string, charsCount byte) *Reflector {
	r := EmptyReflector()

	r.charsCount = charsCount

	r.initWiring(&encoding)

	return r
}

func EmptyReflector() *Reflector {
	return &Reflector{}
}

func (r *Reflector) Forward(c byte) byte {
	return r.wiring[c]
}

func (r *Reflector) initWiring(encoding *string) {
	r.wiring = make([]byte, r.charsCount)

	var i byte

	for ; i < r.charsCount; i++ {
		r.wiring[i] = (*encoding)[i] - UpperA
	}
}