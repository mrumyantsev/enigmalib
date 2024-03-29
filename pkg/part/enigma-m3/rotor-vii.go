package enigmam3

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorVII struct {
	*machine.Rotor
}

func NewRotorVII(pos byte, ring byte) *RotorVII {
	return &RotorVII{
		machine.NewRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "ZM", pos, ring, base.CharactersCount26),
	}
}
