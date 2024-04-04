package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type M4SReflectorBThin struct {
	*base.Reflector
}

func NewM4SReflectorBThin() *M4SReflectorBThin {
	return &M4SReflectorBThin{
		base.NewReflector("ENKQAUYWJICOPBLMDXZVFTHRGS", base.CharsCount26),
	}
}
