package base

type Translator struct {
}

func NewTranslator() *Translator {
	return &Translator{}
}

func EmptyTranslator() *Reflector {
	return &Reflector{}
}

func (t *Translator) TranslateIn(c byte) byte {
	return c - UpperA
}

func (t *Translator) TranslateOut(c byte) byte {
	return c + UpperA
}
