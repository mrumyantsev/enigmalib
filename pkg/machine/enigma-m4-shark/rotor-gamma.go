package enigmam4shark

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type RotorGamma struct {
	*machine.Rotor
}

func NewRotorGamma(pos byte, ring byte) *RotorGamma {
	return &RotorGamma{
		machine.NewRotor("FSOKANUERHMBTIYCWLQPZXVGJD", "", pos, ring, CharactersCount),
	}
}

// @Override
func (r *RotorGamma) Turnover() {
	// No action: this rotor stays still all the time
}
