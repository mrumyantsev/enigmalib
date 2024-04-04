package main

import (
	"fmt"
	"log"

	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
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
			{Name: "BETA", Position: 0, Ring: 0},
			{Name: "I", Position: 0, Ring: 0},
			{Name: "II", Position: 0, Ring: 0},
			{Name: "III", Position: 0, Ring: 0},
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
