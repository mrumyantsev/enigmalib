package machine

import "github.com/mrumyantsev/enigma/pkg/base"

type Translator struct {
}

func NewTranslator() *Translator {
	return &Translator{}
}

func EmptyTranslator() *Reflector {
	return &Reflector{}
}

func (t *Translator) TranslateIn(c byte) byte {
	return c - base.AlphaCharsOffset
}

func (t *Translator) TranslateOut(c byte) byte {
	return c + base.AlphaCharsOffset
}
