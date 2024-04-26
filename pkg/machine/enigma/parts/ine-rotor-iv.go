package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type INERotorIV struct {
	*base.Rotor
}

func NewINERotorIV(pos byte, ring byte) *INERotorIV {
	return &INERotorIV{
		base.NewRotor("FGZJMVXEPBWSHQTLIUDYKCNRAO", "J", pos, ring, base.CharsCount26),
	}
}
