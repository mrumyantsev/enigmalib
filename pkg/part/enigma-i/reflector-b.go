package enigmai

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type ReflectorB struct {
	*machine.Reflector
}

func NewReflectorB() *ReflectorB {
	return &ReflectorB{
		machine.NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT", base.CharactersCount26),
	}
}
