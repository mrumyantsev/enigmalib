package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type IReflectorB struct {
	*base.Reflector
}

func NewIReflectorB() *IReflectorB {
	return &IReflectorB{
		base.NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT", base.CharsCount26),
	}
}
