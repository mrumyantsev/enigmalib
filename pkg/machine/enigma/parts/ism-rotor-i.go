package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type ISMRotorI struct {
	*base.Rotor
}

func NewISMRotorI(pos byte, ring byte) *ISMRotorI {
	return &ISMRotorI{
		base.NewRotor("VEOSIRZUJDQCKGWYPNXAFLTHMB", "Q", pos, ring, base.CharsCount26),
	}
}
