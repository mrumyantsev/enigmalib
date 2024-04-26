package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type M3RotorVI struct {
	*base.Rotor
}

func NewM3RotorVI(pos byte, ring byte) *M3RotorVI {
	return &M3RotorVI{
		base.NewRotor("JPGVOUMFYQBENHZRDKASXLICTW", "ZM", pos, ring, base.CharsCount26),
	}
}
