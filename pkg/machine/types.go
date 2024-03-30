package machine

type Machiner interface {
	EncryptString(msg string) (string, error)
	Encrypt(msg []byte) ([]byte, error)
	EncryptChar(c byte) byte
}

type RotorSettings struct {
	Name    string
	Pos     byte
	RingPos byte
}
