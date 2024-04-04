package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type INEReflector struct {
	*base.Reflector
}

func NewINEReflector() *INEReflector {
	return &INEReflector{
		base.NewReflector("MOWJYPUXNDSRAIBFVLKZGQCHET", base.CharsCount26),
	}
}
