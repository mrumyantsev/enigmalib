package base

type Translator struct {
}

func NewTranslator() *Translator {
	return new(Translator)
}

func (t *Translator) TranslateIn(c byte) byte {
	return c - UpperA
}

func (t *Translator) TranslateOut(c byte) byte {
	return c + UpperA
}
