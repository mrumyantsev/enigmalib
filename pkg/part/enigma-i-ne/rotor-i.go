package enigmaine

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorI struct {
	*machine.Rotor
}

func NewRotorI(pos byte, ring byte) *RotorI {
	return &RotorI{
		machine.NewRotor("WTOKASUYVRBXJHQCPZEFMDINLG", "Q", pos, ring, base.CharactersCount26),
	}
}
