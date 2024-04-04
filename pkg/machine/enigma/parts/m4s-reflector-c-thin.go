package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type M4SReflectorCThin struct {
	*base.Reflector
}

func NewM4SReflectorCThin() *M4SReflectorCThin {
	return &M4SReflectorCThin{
		base.NewReflector("RDOBJNTKVEHMLFCWZAXGYIPSUQ", base.CharsCount26),
	}
}
