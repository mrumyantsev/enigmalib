package enigmai

import "github.com/mrumyantsev/enigma/pkg/machine"

type ReflectorC struct {
	*machine.Reflector
}

func NewReflectorC() *ReflectorC {
	return &ReflectorC{
		machine.NewReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL", CharactersCount),
	}
}
