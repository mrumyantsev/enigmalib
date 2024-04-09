package models

import "github.com/mrumyantsev/cipher-machines-app/pkg/machine"

type PlaintextMsg struct {
	Plaintext string   `json:"plaintext"`
	Machine   string   `json:"machine"`
	Settings  Settings `json:"settings"`
}

type Settings struct {
	Reflector     string                  `json:"reflector"`
	RotorSettings []machine.RotorSettings `json:"rotorSettings"`
	Plugboard     string                  `json:"plugboard"`
}

type CiphertextMsg struct {
	Ciphertext string `json:"ciphertext"`
}
