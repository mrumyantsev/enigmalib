package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i"
	enigmam3 "github.com/mrumyantsev/encryption-app/pkg/part/enigma-m3"
	enigmam4s "github.com/mrumyantsev/encryption-app/pkg/part/enigma-m4-s"
)

type EnigmaM4Shark struct {
	*machine.Machine
}

func NewEnigmaM4Shark(reflector string, rotorSet [base.RotorsCount4]base.RotorSettings, pboard string) *EnigmaM4Shark {
	var refl base.Reflector

	switch reflector {
	case "C THIN":
		refl = enigmam4s.NewReflectorCThin()
	default:
		refl = enigmam4s.NewReflectorBThin()
	}

	rots := make([]base.Rotor, base.RotorsCount4)
	var settings base.RotorSettings

	for i := 0; i < base.RotorsCount4; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = enigmai.NewRotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = enigmai.NewRotorIII(settings.Pos, settings.RingPos)
		case "IV":
			rots[i] = enigmai.NewRotorIV(settings.Pos, settings.RingPos)
		case "V":
			rots[i] = enigmai.NewRotorV(settings.Pos, settings.RingPos)
		case "VI":
			rots[i] = enigmam3.NewRotorVI(settings.Pos, settings.RingPos)
		case "VII":
			rots[i] = enigmam3.NewRotorVII(settings.Pos, settings.RingPos)
		case "VIII":
			rots[i] = enigmam3.NewRotorVIII(settings.Pos, settings.RingPos)
		case "BETA":
			rots[i] = enigmam4s.NewRotorBeta(settings.Pos, settings.RingPos)
		case "GAMMA":
			rots[i] = enigmam4s.NewRotorGamma(settings.Pos, settings.RingPos)
		default:
			rots[i] = enigmai.NewRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := enigmai.NewStator()

	pb := machine.NewPlugboard(pboard, base.CharactersCount26)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaM4Shark{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
