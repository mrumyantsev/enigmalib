package machine

import (
	"errors"

	"github.com/mrumyantsev/enigma/pkg/base"
)

var (
	errReflectorNotSet = errors.New("reflector not set")
)

type Machine struct {
	// Inner parts
	reflector base.Reflector
	rotors    []base.Rotor
	stator    base.Stator
	plugboard base.Plugboard

	// Programmable attachments
	filter     base.Filter
	translator base.Translator
}

func NewMachine(refl base.Reflector, rots []base.Rotor, sta base.Stator, pb base.Plugboard, fil base.Filter, tr base.Translator) *Machine {
	return &Machine{
		reflector: refl,
		rotors:    rots,
		stator:    sta,
		plugboard: pb,

		filter:     fil,
		translator: tr,
	}
}

func (m *Machine) EncryptString(msg string) (string, error) {
	res, err := m.Encrypt([]byte(msg))
	if err != nil {
		return "", err
	}

	return string(res), err
}

func (m *Machine) Encrypt(msg []byte) ([]byte, error) {
	if m.reflector == nil {
		return nil, errReflectorNotSet
	}

	msgLen := len(msg)
	res := make([]byte, 0, msgLen)
	var c byte

	for i := 0; i < msgLen; i++ {
		c = m.EncryptChar(msg[i])

		if c == 0 {
			continue
		}

		res = append(res, c)
	}

	return res, nil
}

func (m *Machine) EncryptChar(c byte) byte {
	c = m.filter.Filter(c)

	if c == 0 {
		return 0
	}

	c = m.translator.TranslateIn(c)

	m.rotateRotors()

	c = m.plugboard.Forward(c)

	c = m.stator.Forward(c)

	rotLen := len(m.rotors)

	for i := 0; i < rotLen; i++ {
		c = m.rotors[i].Forward(c)
	}

	c = m.reflector.Forward(c)

	for i := rotLen - 1; i >= 0; i-- {
		c = m.rotors[i].Backward(c)
	}

	c = m.stator.Backward(c)

	c = m.plugboard.Forward(c)

	c = m.translator.TranslateOut(c)

	return c
}

func (m *Machine) rotateRotors() {
	rotLen := len(m.rotors)

	if rotLen == 0 {
		return
	}

	turnovers := 1

	for i := 0; i < rotLen; i++ {
		if m.rotors[i].IsAtNotch() {
			turnovers++

			continue
		}

		break
	}

	if turnovers > rotLen {
		turnovers = rotLen
	}

	for i := 0; i < turnovers; i++ {
		m.rotors[i].Turnover()
	}
}
