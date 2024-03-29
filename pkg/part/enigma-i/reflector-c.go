package enigmai

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type ReflectorC struct {
	*machine.Reflector
}

func NewReflectorC() *ReflectorC {
	return &ReflectorC{
		machine.NewReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL", base.CharactersCount26),
	}
}
