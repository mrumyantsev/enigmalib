package enigmam4s

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorBeta struct {
	*machine.Rotor
}

func NewRotorBeta(pos byte, ring byte) *RotorBeta {
	return &RotorBeta{
		machine.NewRotor("LEYJVCNIXWPBQMDRTAKZGFUHOS", "", pos, ring, base.CharactersCount26),
	}
}

// @Override
func (r *RotorBeta) Turnover() {
	// No action: this rotor stays still all the time
}
