package enigmam3

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type RotorVI struct {
	*machine.Rotor
}

func NewRotorVI(pos byte, ring byte) *RotorVI {
	return &RotorVI{
		machine.NewRotor("JPGVOUMFYQBENHZRDKASXLICTW", "ZM", pos, ring, CharactersCount),
	}
}
