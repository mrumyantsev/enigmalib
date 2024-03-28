package enigmam4shark

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type ReflectorCThin struct {
	*machine.Reflector
}

func NewReflectorCThin() *ReflectorCThin {
	return &ReflectorCThin{
		machine.NewReflector("RDOBJNTKVEHMLFCWZAXGYIPSUQ", CharactersCount),
	}
}
