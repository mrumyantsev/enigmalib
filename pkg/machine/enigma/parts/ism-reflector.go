package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type ISMReflector struct {
	*base.Reflector
}

func NewISMReflector() *ISMReflector {
	return &ISMReflector{
		base.NewReflector("CIAGSNDRBYTPZFULVHEKOQXWJM", base.CharsCount26),
	}
}
