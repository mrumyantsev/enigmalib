package enigmainorenigma

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type RotorIV struct {
	*machine.Rotor
}

func NewRotorIV(pos byte, ring byte) *RotorIV {
	return &RotorIV{
		machine.NewRotor("FGZJMVXEPBWSHQTLIUDYKCNRAO", "J", pos, ring, CharactersCount),
	}
}
