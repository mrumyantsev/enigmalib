package machine

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
)

type Rotor struct {
	*Stator
	notches []bool
	pos     byte
	ring    byte
}

func NewRotor(encoding string, notchPoses string, pos byte, ring byte, charCount byte) *Rotor {
	r := EmptyRotor()

	r.charCount = charCount

	r.pos = pos

	r.ring = ring

	r.initWiring(&encoding)

	r.initInverseWiring(&encoding)

	r.initNotches(&notchPoses)

	return r
}

func EmptyRotor() *Rotor {
	return &Rotor{Stator: EmptyStator()}
}

func (r *Rotor) Forward(c byte) byte {
	return r.remap(c, r.wiring)
}

func (r *Rotor) Backward(c byte) byte {
	return r.remap(c, r.inverse)
}

func (r *Rotor) Turnover() {
	r.pos = (r.pos + 1) % r.charCount
}

func (r *Rotor) IsAtNotch() bool {
	return r.notches[r.pos]
}

func (r *Rotor) initNotches(notchPoses *string) {
	r.notches = make([]bool, r.charCount)
	notchPosesLen := len(*notchPoses)

	for i := 0; i < notchPosesLen; i++ {
		r.notches[(*notchPoses)[i]-base.AlphaCharsOffset] = true
	}
}

func (r *Rotor) remap(c byte, wiring []byte) byte {
	shift := r.pos - r.ring

	return (wiring[(c+shift+r.charCount)%r.charCount] - shift + r.charCount) % r.charCount
}
