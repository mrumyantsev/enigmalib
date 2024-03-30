package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type INERotorIII struct {
	*base.Rotor
}

func NewINERotorIII(pos byte, ring byte) *INERotorIII {
	return &INERotorIII{
		base.NewRotor("JWFMHNBPUSDYTIXVZGRQLAOEKC", "V", pos, ring, base.CharsCount26),
	}
}
