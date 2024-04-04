package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type IReflectorA struct {
	*base.Reflector
}

func NewIReflectorA() *IReflectorA {
	return &IReflectorA{
		base.NewReflector("EJMZALYXVBWFCRQUONTSPIKHGD", base.CharsCount26),
	}
}
