package base

import (
	"strings"
)

type Plugboard struct {
	*Reflector
}

func NewPlugboard(plugs string, charsCount byte) *Plugboard {
	p := EmptyPlugboard()

	p.charsCount = charsCount

	p.initWiring(&plugs)

	return p
}

func EmptyPlugboard() *Plugboard {
	return &Plugboard{Reflector: EmptyReflector()}
}

func (p *Plugboard) initWiring(plugs *string) {
	p.wiring = make([]byte, p.charsCount)

	var i byte

	for ; i < p.charsCount; i++ {
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
		c1 = changes[i][0] - UpperA
		c2 = changes[i][1] - UpperA

		p.wiring[c1] = c2
		p.wiring[c2] = c1
	}
}
