func main() {
	flag := true
	for flag {
		MenuInicio := "Rebeldes\n	[ 1 ] Consultar Datos\n	[ 2 ] Cerrar programa\nIngrese su opción:"
		MenuConsulta := "Consulta\n	[ 1 ] Consultar Datos de Logística\n	[ 2 ] Consultar Datos Financieros\n	[ 3 ] Consultar Datos Militares\n	[ 4 ] Volver\nIngrese su opción:"
		fmt.Print(MenuInicio)

		var eleccion string
		fmt.Scanln(&eleccion)

		switch eleccion {
		case "1":

			fmt.Print(MenuConsulta)
			var eleccionC string
			fmt.Scanln(&eleccionC)
			switch eleccionC {
			case "1":
				fmt.Println("Entregando Datos de Logistica")

			case "2":
				fmt.Println("Entregando Datos Financieron")

			case "3":
				fmt.Println("Entregando Datos Militares")

			case "4":
				fmt.Println("Volviendo")
			default:
				fmt.Println("Opcion no valida, volviendo al menu de inicio")
			}

		case "2":

			fmt.Println("Realizar ACA cierre de las otras cosas")
			flag = false

		default:
			fmt.Println("Opcion no valida")
		}
	}
	fmt.Println("Se acabo este programa")

}
