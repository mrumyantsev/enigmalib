package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type IRotorII struct {
	*base.Rotor
}

func NewIRotorII(pos byte, ring byte) *IRotorII {
	return &IRotorII{
		base.NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "E", pos, ring, base.CharsCount26),
	}
}
