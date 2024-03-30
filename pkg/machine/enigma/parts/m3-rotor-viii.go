package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type M3RotorVIII struct {
	*base.Rotor
}

func NewM3RotorVIII(pos byte, ring byte) *M3RotorVIII {
	return &M3RotorVIII{
		base.NewRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "ZM", pos, ring, base.CharsCount26),
	}
}
