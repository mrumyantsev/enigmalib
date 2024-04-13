package machine

type Machiner interface {
	EncryptString(msg string) (string, error)
	Encrypt(msg []byte) ([]byte, error)
	EncryptChar(c byte) byte
}

type RotorSettings struct {
	Name     string `json:"name"`
	Position byte   `json:"position"`
	Ring     byte   `json:"ring"`
}
