package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type IRotorIII struct {
	*base.Rotor
}

func NewIRotorIII(pos byte, ring byte) *IRotorIII {
	return &IRotorIII{
		base.NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "V", pos, ring, base.CharsCount26),
	}
}
