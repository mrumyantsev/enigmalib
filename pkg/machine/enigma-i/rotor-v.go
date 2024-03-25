package enigmai

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorV struct {
	*machine.Rotor
}

func NewRotorV(pos byte, ring byte) *RotorV {
	return &RotorV{
		machine.NewRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "Z", pos, ring, CharactersCount),
	}
}
