package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type ISMRotorII struct {
	*base.Rotor
}

func NewISMRotorII(pos byte, ring byte) *ISMRotorII {
	return &ISMRotorII{
		base.NewRotor("UEMOATQLSHPKCYFWJZBGVXIDNR", "E", pos, ring, base.CharsCount26),
	}
}
