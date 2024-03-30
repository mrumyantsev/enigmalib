package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type INEReflector struct {
	*base.Reflector
}

func NewINEReflector() *INEReflector {
	return &INEReflector{
		base.NewReflector("MOWJYPUXNDSRAIBFVLKZGQCHET", base.CharsCount26),
	}
}
