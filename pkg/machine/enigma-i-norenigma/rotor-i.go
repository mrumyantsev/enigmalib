package enigmainorenigma

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorI struct {
	*machine.Rotor
}

func NewRotorI(pos byte, ring byte) *RotorI {
	return &RotorI{
		machine.NewRotor("WTOKASUYVRBXJHQCPZEFMDINLG", "Q", pos, ring, CharactersCount),
	}
}
