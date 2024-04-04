package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type ISMRotorIII struct {
	*base.Rotor
}

func NewISMRotorIII(pos byte, ring byte) *ISMRotorIII {
	return &ISMRotorIII{
		base.NewRotor("TZHXMBSIPNURJFDKEQVCWGLAOY", "V", pos, ring, base.CharsCount26),
	}
}
