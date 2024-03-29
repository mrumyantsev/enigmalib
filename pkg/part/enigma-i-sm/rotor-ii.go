package enigmaism

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorII struct {
	*machine.Rotor
}

func NewRotorII(pos byte, ring byte) *RotorII {
	return &RotorII{
		machine.NewRotor("UEMOATQLSHPKCYFWJZBGVXIDNR", "E", pos, ring, base.CharactersCount26),
	}
}
