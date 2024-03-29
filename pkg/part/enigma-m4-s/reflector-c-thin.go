package enigmam4s

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type ReflectorCThin struct {
	*machine.Reflector
}

func NewReflectorCThin() *ReflectorCThin {
	return &ReflectorCThin{
		machine.NewReflector("RDOBJNTKVEHMLFCWZAXGYIPSUQ", base.CharactersCount26),
	}
}
