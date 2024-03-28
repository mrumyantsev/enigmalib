package enigmai

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type RotorII struct {
	*machine.Rotor
}

func NewRotorII(pos byte, ring byte) *RotorII {
	return &RotorII{
		machine.NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "E", pos, ring, CharactersCount),
	}
}
