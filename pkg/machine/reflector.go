package machine

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
)

type Reflector struct {
	wiring    []byte
	charCount byte
}

func NewReflector(encoding string, charCount byte) *Reflector {
	r := EmptyReflector()

	r.charCount = charCount

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
	r.wiring = make([]byte, r.charCount)

	var i byte

	for ; i < r.charCount; i++ {
		r.wiring[i] = (*encoding)[i] - base.AlphaCharsOffset
	}
}
