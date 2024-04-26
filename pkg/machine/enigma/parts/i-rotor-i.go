package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type IRotorI struct {
	*base.Rotor
}

func NewIRotorI(pos byte, ring byte) *IRotorI {
	return &IRotorI{
		base.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q", pos, ring, base.CharsCount26),
	}
}
