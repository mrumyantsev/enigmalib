package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type INERotorI struct {
	*base.Rotor
}

func NewINERotorI(pos byte, ring byte) *INERotorI {
	return &INERotorI{
		base.NewRotor("WTOKASUYVRBXJHQCPZEFMDINLG", "Q", pos, ring, base.CharsCount26),
	}
}
