package parts

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

type IStator struct {
	*base.Stator
}

func NewIStator() *IStator {
	return &IStator{
		base.NewStator("ABCDEFGHIJKLMNOPQRSTUVWXYZ", base.CharsCount26),
	}
}
