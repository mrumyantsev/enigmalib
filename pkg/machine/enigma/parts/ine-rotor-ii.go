package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type INERotorII struct {
	*base.Rotor
}

func NewINERotorII(pos byte, ring byte) *INERotorII {
	return &INERotorII{
		base.NewRotor("GJLPUBSWEMCTQVHXAOFZDRKYNI", "E", pos, ring, base.CharsCount26),
	}
}
