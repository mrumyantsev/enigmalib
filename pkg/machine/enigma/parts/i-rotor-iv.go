package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type IRotorIV struct {
	*base.Rotor
}

func NewIRotorIV(pos byte, ring byte) *IRotorIV {
	return &IRotorIV{
		base.NewRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "J", pos, ring, base.CharsCount26),
	}
}
