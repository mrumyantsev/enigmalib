package enigmaine

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorII struct {
	*machine.Rotor
}

func NewRotorII(pos byte, ring byte) *RotorII {
	return &RotorII{
		machine.NewRotor("GJLPUBSWEMCTQVHXAOFZDRKYNI", "E", pos, ring, base.CharactersCount26),
	}
}
