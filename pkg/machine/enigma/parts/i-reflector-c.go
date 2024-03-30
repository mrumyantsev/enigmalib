package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type IReflectorC struct {
	*base.Reflector
}

func NewIReflectorC() *IReflectorC {
	return &IReflectorC{
		base.NewReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL", base.CharsCount26),
	}
}
