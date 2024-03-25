package enigmainorenigma

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorII struct {
	*machine.Rotor
}

func NewRotorII(pos byte, ring byte) *RotorII {
	return &RotorII{
		machine.NewRotor("GJLPUBSWEMCTQVHXAOFZDRKYNI", "E", pos, ring, CharactersCount),
	}
}
