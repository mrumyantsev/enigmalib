package enigmaisondermaschine

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type Reflector struct {
	*machine.Reflector
}

func NewReflector() *Reflector {
	return &Reflector{
		machine.NewReflector("CIAGSNDRBYTPZFULVHEKOQXWJM", CharactersCount),
	}
}
