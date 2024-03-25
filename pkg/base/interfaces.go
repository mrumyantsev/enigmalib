package base

type Machine interface {
	EncryptString(msg string) (string, error)
	Encrypt(msg []byte) ([]byte, error)
	EncryptChar(c byte) byte
}

type Reflector interface {
	Forward(c byte) byte
}

type Rotor interface {
	Forward(c byte) byte
	Backward(c byte) byte
	Turnover()
	IsAtNotch() bool
}

type Stator interface {
	Forward(c byte) byte
	Backward(c byte) byte
}

type Plugboard interface {
	Forward(c byte) byte
}

type Filter interface {
	Filter(c byte) byte
}

type Translator interface {
	TranslateIn(c byte) byte
	TranslateOut(c byte) byte
}
