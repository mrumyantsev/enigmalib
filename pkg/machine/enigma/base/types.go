package base

type Reflectorer interface {
	Forward(c byte) byte
}

type Rotorer interface {
	Forward(c byte) byte
	Backward(c byte) byte
	Turnover()
	IsAtNotch() bool
}

type Statorer interface {
	Forward(c byte) byte
	Backward(c byte) byte
}

type Plugboarder interface {
	Forward(c byte) byte
}

type Filterer interface {
	Filter(c byte) byte
}

type Translatorer interface {
	TranslateIn(c byte) byte
	TranslateOut(c byte) byte
}
