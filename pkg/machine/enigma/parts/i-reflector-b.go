package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type IReflectorB struct {
	*base.Reflector
}

func NewIReflectorB() *IReflectorB {
	return &IReflectorB{
		base.NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT", base.CharsCount26),
	}
}
