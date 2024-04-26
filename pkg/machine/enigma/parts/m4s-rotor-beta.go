package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type M4SRotorBeta struct {
	*base.Rotor
}

func NewM4SRotorBeta(pos byte, ring byte) *M4SRotorBeta {
	return &M4SRotorBeta{
		base.NewRotor("LEYJVCNIXWPBQMDRTAKZGFUHOS", "", pos, ring, base.CharsCount26),
	}
}

// @Override
func (r *M4SRotorBeta) Turnover() {
	// No action: this rotor stays still all the time
}
