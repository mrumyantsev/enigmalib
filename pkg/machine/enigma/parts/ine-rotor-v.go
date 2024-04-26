package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type INERotorV struct {
	*base.Rotor
}

func NewINERotorV(pos byte, ring byte) *INERotorV {
	return &INERotorV{
		base.NewRotor("HEJXQOTZBVFDASCILWPGYNMURK", "Z", pos, ring, base.CharsCount26),
	}
}
