package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type M3RotorVII struct {
	*base.Rotor
}

func NewM3RotorVII(pos byte, ring byte) *M3RotorVII {
	return &M3RotorVII{
		base.NewRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "ZM", pos, ring, base.CharsCount26),
	}
}
