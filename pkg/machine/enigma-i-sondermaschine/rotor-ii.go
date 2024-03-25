package enigmaisondermaschine

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorII struct {
	*machine.Rotor
}

func NewRotorII(pos byte, ring byte) *RotorII {
	return &RotorII{
		machine.NewRotor("UEMOATQLSHPKCYFWJZBGVXIDNR", "E", pos, ring, CharactersCount),
	}
}
