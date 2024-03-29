package enigmai

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorIV struct {
	*machine.Rotor
}

func NewRotorIV(pos byte, ring byte) *RotorIV {
	return &RotorIV{
		machine.NewRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "J", pos, ring, base.CharactersCount26),
	}
}
