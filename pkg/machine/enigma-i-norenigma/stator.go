package enigmainorenigma

import "github.com/mrumyantsev/enigma/pkg/machine"

type Stator struct {
	*machine.Stator
}

func NewStator() *Stator {
	return &Stator{
		machine.NewStator("ABCDEFGHIJKLMNOPQRSTUVWXYZ", CharactersCount),
	}
}
