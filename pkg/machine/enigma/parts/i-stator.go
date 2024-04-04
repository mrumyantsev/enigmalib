package parts

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type IStator struct {
	*base.Stator
}

func NewIStator() *IStator {
	return &IStator{
		base.NewStator("ABCDEFGHIJKLMNOPQRSTUVWXYZ", base.CharsCount26),
	}
}
