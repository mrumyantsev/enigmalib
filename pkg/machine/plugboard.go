package machine

import (
	"strings"

	"github.com/mrumyantsev/enigma/pkg/base"
)

type Plugboard struct {
	*Reflector
}

func NewPlugboard(plugs string, charCount byte) *Plugboard {
	p := EmptyPlugboard()

	p.charCount = charCount

	p.initWiring(&plugs)

	return p
}

func EmptyPlugboard() *Plugboard {
	return &Plugboard{Reflector: EmptyReflector()}
}

func (p *Plugboard) initWiring(plugs *string) {
	p.wiring = make([]byte, p.charCount)

	var i byte

	for ; i < p.charCount; i++ {
		p.wiring[i] = i
	}

	if *plugs == "" {
		return
	}

	changes := strings.Split(*plugs, " ")
	changesLen := len(changes)

	var c1 byte
	var c2 byte

	for i := 0; i < changesLen; i++ {
		c1 = changes[i][0] - base.AlphaCharsOffset
		c2 = changes[i][1] - base.AlphaCharsOffset

		p.wiring[c1] = c2
		p.wiring[c2] = c1
	}
}
