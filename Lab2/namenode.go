package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DateNodeRandom() (Nombre_DateNode string, IP string) {
	rand.Seed(time.Now().UnixNano())
	switch os := rand.Intn(3); os {
	case 0:
		Nombre := "DateNode Grunt"
		IP := "IP DATENODE"
		return Nombre, IP
	case 1:
		Nombre := "DateNode Synth"
		IP := "IP DATENODE"
		return Nombre, IP
	default:
		Nombre := "DateNode Cremator"
		IP := "IP DATENODE"
		return Nombre, IP
	}
}


