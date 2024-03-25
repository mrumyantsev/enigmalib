package enigmainorenigma

import "github.com/mrumyantsev/enigma/pkg/machine"

type RotorIII struct {
	*machine.Rotor
}

func NewRotorIII(pos byte, ring byte) *RotorIII {
	return &RotorIII{
		machine.NewRotor("JWFMHNBPUSDYTIXVZGRQLAOEKC", "V", pos, ring, CharactersCount),
	}
}
