package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type IReflectorC struct {
	*base.Reflector
}

func NewIReflectorC() *IReflectorC {
	return &IReflectorC{
		base.NewReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL", base.CharsCount26),
	}
}
