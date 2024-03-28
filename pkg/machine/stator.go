package machine

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
)

type Stator struct {
	*Reflector
	inverse []byte
}

func NewStator(encoding string, charCount byte) *Stator {
	s := EmptyStator()

	s.charCount = charCount

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
	s.inverse = make([]byte, s.charCount)

	var i byte

	for ; i < s.charCount; i++ {
		s.inverse[(*encoding)[i]-base.AlphaCharsOffset] = i
	}
}
