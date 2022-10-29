package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
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

func RevisarID(ID string) bool {
	file, err := os.Open("DATA.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		Split_Msj := strings.Split(scanner.Text(), ":")
		if Split_Msj[1] == ID {

			file.Close()
			return false

		}
	}

	file.Close()
	return true

}

func RetornarData(Tipo string) string {
	file, err := os.Open("DATA.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	StringRetorno := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		Split_Msj := strings.Split(scanner.Text(), ":")

		if Split_Msj[0] == Tipo {

			StringRetorno = StringRetorno + Split_Msj[1] + ":" + Split_Msj[2] + "\n"

		}
	}

	file.Close()
	return StringRetorno
}

func main() {

	file, err := os.Create("DATA.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	string1 := "MILITAR:112:LLEGADA PISTOLAS"

	Split_Msj := strings.Split(string1, ":")
	Tipo := Split_Msj[0]
	ID := Split_Msj[1]

	NAMEDATENODE, _ := DateNodeRandom()

	file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	string1 = "LOGISTICA:234:DATA GENERICA"

	Split_Msj = strings.Split(string1, ":")
	Tipo = Split_Msj[0]
	ID = Split_Msj[1]

	NAMEDATENODE, _ = DateNodeRandom()

	file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	string1 = "FINANCIEra:32134:DATA GENERICA"

	Split_Msj = strings.Split(string1, ":")
	Tipo = Split_Msj[0]
	ID = Split_Msj[1]

	NAMEDATENODE, _ = DateNodeRandom()

	file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	string1 = "LOGISTICA:2334:DATA GENERICA"

	Split_Msj = strings.Split(string1, ":")
	Tipo = Split_Msj[0]
	ID = Split_Msj[1]

	NAMEDATENODE, _ = DateNodeRandom()

	file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	string1 = "MILITAR:23400:DATA GENERICA"

	Split_Msj = strings.Split(string1, ":")
	Tipo = Split_Msj[0]
	ID = Split_Msj[1]

	NAMEDATENODE, _ = DateNodeRandom()

	file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	string1 = "FINANCIERA:234000:DATA GENERICA"

	Split_Msj = strings.Split(string1, ":")
	Tipo = Split_Msj[0]
	ID = Split_Msj[1]

	NAMEDATENODE, _ = DateNodeRandom()

	file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	scanner := bufio.NewScanner(file)
	println("LLEGo")
	for scanner.Scan() {
		println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

}