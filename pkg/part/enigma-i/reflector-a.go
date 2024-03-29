package enigmai

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type ReflectorA struct {
	*machine.Reflector
}

func NewReflectorA() *ReflectorA {
	return &ReflectorA{
		machine.NewReflector("EJMZALYXVBWFCRQUONTSPIKHGD", base.CharactersCount26),
	}
}
