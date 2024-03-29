package enigmam4s

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type RotorGamma struct {
	*machine.Rotor
}

func NewRotorGamma(pos byte, ring byte) *RotorGamma {
	return &RotorGamma{
		machine.NewRotor("FSOKANUERHMBTIYCWLQPZXVGJD", "", pos, ring, base.CharactersCount26),
	}
}

// @Override
func (r *RotorGamma) Turnover() {
	// No action: this rotor stays still all the time
}
