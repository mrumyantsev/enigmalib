package enigmam3

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/machine/enigma-i"
)

const (
	CharactersCount = 26
	RotorsCount     = 3
)

type EnigmaM3 struct {
	*machine.Machine
}

func NewEnigmaM3(reflector string, rotorSet [RotorsCount]base.RotorSettings, stator string, pboard string) *EnigmaM3 {
	var refl base.Reflector

	switch reflector {
	case "C":
		refl = enigmai.NewReflectorC()
	default:
		refl = enigmai.NewReflectorB()
	}

	rots := make([]base.Rotor, RotorsCount)

	for i := 0; i < RotorsCount; i++ {
		switch rotorSet[i].Name {
		case "II":
			rots[i] = enigmai.NewRotorII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "III":
			rots[i] = enigmai.NewRotorIII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "IV":
			rots[i] = enigmai.NewRotorIV(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "V":
			rots[i] = enigmai.NewRotorV(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "VI":
			rots[i] = NewRotorVI(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "VII":
			rots[i] = NewRotorVII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "VIII":
			rots[i] = NewRotorVIII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		default:
			rots[i] = enigmai.NewRotorI(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		}
	}

	var sta base.Stator

	switch stator {
	default:
		sta = enigmai.NewStator()
	}

	pb := machine.NewPlugboard(pboard, CharactersCount)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaM3{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
