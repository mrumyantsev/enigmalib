package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type IRotorV struct {
	*base.Rotor
}

func NewIRotorV(pos byte, ring byte) *IRotorV {
	return &IRotorV{
		base.NewRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "Z", pos, ring, base.CharsCount26),
	}
}
