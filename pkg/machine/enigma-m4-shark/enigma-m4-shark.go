package enigmam4shark

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/machine/enigma-i"
	enigmam3 "github.com/mrumyantsev/encryption-app/pkg/machine/enigma-m3"
)

const (
	CharactersCount = 26
	RotorsCount     = 4
)

type EnigmaM4Shark struct {
	*machine.Machine
}

func NewEnigmaM4Shark(reflector string, rotorSet [RotorsCount]base.RotorSettings, stator string, pboard string) *EnigmaM4Shark {
	var refl base.Reflector

	switch reflector {
	case "C THIN":
		refl = NewReflectorCThin()
	default:
		refl = NewReflectorBThin()
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
			rots[i] = enigmam3.NewRotorVI(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "VII":
			rots[i] = enigmam3.NewRotorVII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "VIII":
			rots[i] = enigmam3.NewRotorVIII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "BETA":
			rots[i] = NewRotorBeta(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "GAMMA":
			rots[i] = NewRotorGamma(
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

	return &EnigmaM4Shark{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
