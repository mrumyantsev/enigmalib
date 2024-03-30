package base

type Stator struct {
	*Reflector
	inverse []byte
}

func NewStator(encoding string, charsCount byte) *Stator {
	s := EmptyStator()

	s.charsCount = charsCount

	s.initWiring(&encoding)

	s.initInverseWiring(&encoding)

	return s
}

func EmptyStator() *Stator {
	return &Stator{Reflector: EmptyReflector()}
}

func (s *Stator) Backward(c byte) byte {
	return s.inverse[c]
}

func (s *Stator) initInverseWiring(encoding *string) {
	s.inverse = make([]byte, s.charsCount)

	var i byte

	for ; i < s.charsCount; i++ {
		s.inverse[(*encoding)[i]-UpperA] = i
	}
}