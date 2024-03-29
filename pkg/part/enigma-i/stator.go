package enigmai

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

type Stator struct {
	*machine.Stator
}

func NewStator() *Stator {
	return &Stator{
		machine.NewStator("ABCDEFGHIJKLMNOPQRSTUVWXYZ", base.CharactersCount26),
	}
}
