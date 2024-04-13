package base

const (
	CharsCount26 byte = 26

	ReflectorsCount1 = 1
	RotorsCount3     = 3
	RotorsCount4     = 4
	PlugboardsCount1 = 1

	UpperA byte = 65
	UpperZ byte = 90
	LowerA byte = 97
	LowerZ byte = 122
)

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
