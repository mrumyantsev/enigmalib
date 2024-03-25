package enigmai

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorIV struct {
	*machine.Rotor
}

func NewRotorIV(pos byte, ring byte) *RotorIV {
	return &RotorIV{
		machine.NewRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "J", pos, ring, CharactersCount),
	}
}
