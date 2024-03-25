package enigmainorenigma

import "github.com/mrumyantsev/enigma/pkg/machine"

type Reflector struct {
	*machine.Reflector
}

func NewReflector() *Reflector {
	return &Reflector{
		machine.NewReflector("MOWJYPUXNDSRAIBFVLKZGQCHET", CharactersCount),
	}
}
