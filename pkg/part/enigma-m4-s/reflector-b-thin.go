package enigmam4s

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type ReflectorBThin struct {
	*machine.Reflector
}

func NewReflectorBThin() *ReflectorBThin {
	return &ReflectorBThin{
		machine.NewReflector("ENKQAUYWJICOPBLMDXZVFTHRGS", base.CharactersCount26),
	}
}
