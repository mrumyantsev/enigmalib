package enigmai

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorIII struct {
	*machine.Rotor
}

func NewRotorIII(pos byte, ring byte) *RotorIII {
	return &RotorIII{
		machine.NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "V", pos, ring, CharactersCount),
	}
}
