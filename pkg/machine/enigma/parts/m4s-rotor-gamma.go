package parts

import (
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
)

type M4SRotorGamma struct {
	*base.Rotor
}

func NewM4SRotorGamma(pos byte, ring byte) *M4SRotorGamma {
	return &M4SRotorGamma{
		base.NewRotor("FSOKANUERHMBTIYCWLQPZXVGJD", "", pos, ring, base.CharsCount26),
	}
}

// @Override
func (r *M4SRotorGamma) Turnover() {
	// No action: this rotor stays still all the time
}
