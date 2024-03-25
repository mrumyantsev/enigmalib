package enigmaisondermaschine

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorI struct {
	*machine.Rotor
}

func NewRotorI(pos byte, ring byte) *RotorI {
	return &RotorI{
		machine.NewRotor("VEOSIRZUJDQCKGWYPNXAFLTHMB", "Q", pos, ring, CharactersCount),
	}
}
