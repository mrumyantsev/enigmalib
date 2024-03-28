package enigmainorenigma

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type RotorV struct {
	*machine.Rotor
}

func NewRotorV(pos byte, ring byte) *RotorV {
	return &RotorV{
		machine.NewRotor("HEJXQOTZBVFDASCILWPGYNMURK", "Z", pos, ring, CharactersCount),
	}
}
