package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type ISMReflector struct {
	*base.Reflector
}

func NewISMReflector() *ISMReflector {
	return &ISMReflector{
		base.NewReflector("CIAGSNDRBYTPZFULVHEKOQXWJM", base.CharsCount26),
	}
}
