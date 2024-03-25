package enigmai

import "github.com/mrumyantsev/enigma/pkg/machine"

type ReflectorB struct {
	*machine.Reflector
}

func NewReflectorB() *ReflectorB {
	return &ReflectorB{
		machine.NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT", CharactersCount),
	}
}
