package enigmai

import "github.com/mrumyantsev/enigma/pkg/machine"

type ReflectorA struct {
	*machine.Reflector
}

func NewReflectorA() *ReflectorA {
	return &ReflectorA{
		machine.NewReflector("EJMZALYXVBWFCRQUONTSPIKHGD", CharactersCount),
	}
}
