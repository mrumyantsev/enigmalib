package enigmam3

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorVIII struct {
	*machine.Rotor
}

func NewRotorVIII(pos byte, ring byte) *RotorVIII {
	return &RotorVIII{
		machine.NewRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "ZM", pos, ring, base.CharactersCount26),
	}
}
