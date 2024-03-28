package enigmam4shark

import "github.com/mrumyantsev/encryption-app/pkg/machine"

type ReflectorBThin struct {
	*machine.Reflector
}

func NewReflectorBThin() *ReflectorBThin {
	return &ReflectorBThin{
		machine.NewReflector("ENKQAUYWJICOPBLMDXZVFTHRGS", CharactersCount),
	}
}
