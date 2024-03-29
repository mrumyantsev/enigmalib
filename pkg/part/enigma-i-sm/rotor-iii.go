package enigmaism

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorIII struct {
	*machine.Rotor
}

func NewRotorIII(pos byte, ring byte) *RotorIII {
	return &RotorIII{
		machine.NewRotor("TZHXMBSIPNURJFDKEQVCWGLAOY", "V", pos, ring, base.CharactersCount26),
	}
}
