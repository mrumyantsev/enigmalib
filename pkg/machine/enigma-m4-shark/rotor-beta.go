package enigmam4shark

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type RotorBeta struct {
	*machine.Rotor
}

func NewRotorBeta(pos byte, ring byte) *RotorBeta {
	return &RotorBeta{
		machine.NewRotor("LEYJVCNIXWPBQMDRTAKZGFUHOS", "", pos, ring, CharactersCount),
	}
}

// @Override
func (r *RotorBeta) Turnover() {
	// No action: this rotor stays still all the time
}
