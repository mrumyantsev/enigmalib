package main

import (
	"fmt"
	"log"

	"github.com/mrumyantsev/encryption-app/pkg/machine"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
)

func main() {
	machine1 := initMachine()
	machine2 := initMachine()
	filter := base.NewFilter()

	plaintext := "Keine besonderen Ereignisse"

	print(plaintext, filter)

	ciphertext, err := machine1.EncryptString(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	print(ciphertext, nil)

	out, err := machine2.EncryptString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	print(out, nil)
}

func initMachine() machine.Machiner {
	return enigma.NewEnigmaM4S(
		"B THIN",
		[base.RotorsCount4]machine.RotorSettings{
			{Name: "BETA", Pos: 0, RingPos: 0},
			{Name: "I", Pos: 0, RingPos: 0},
			{Name: "II", Pos: 0, RingPos: 0},
			{Name: "III", Pos: 0, RingPos: 0},
		},
		"",
	)
}

func print(msg string, fil base.Filterer) {
	msgLen := len(msg)

	if msgLen == 0 {
		return
	}

	count := 1
	var c byte

	for i := 0; i < msgLen; i++ {
		if fil != nil {
			c = fil.Filter(msg[i])

			if c == 0 {
				continue
			}
		} else {
			c = msg[i]
		}

		fmt.Printf("%c ", c)

		if (count % 5) == 0 {
			fmt.Printf(" ")
		}

		count++
	}

	fmt.Println()
	fmt.Println()
}
