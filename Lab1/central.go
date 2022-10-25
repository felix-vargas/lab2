package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall" //Para manejar el comando ctrl + c

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

// Variable global para determinar que escuadron se va y cual vuelve
var escuadrones = 2

// Funcion que se encarga de entablar conexion sincrona con los laboratorios al leer el primer mensaje de la cola
// de RabbitMQ, toma como input un amqp.Delivery y un archivo para escribir sobre este, esta funcion se separa
// del main con el proposito de poder ejecutarla como una rutina go y poder tener las dos necesarias para ambos escuadrones
func todo(d amqp.Delivery, ff *os.File) {
	port := ":50052"                                          //puerto de la conexion con el laboratorio
	fmt.Println("Pedido de ayuda de " + string(d.Body))       //obtiene el primer mensaje de la cola
	connS, err := grpc.Dial(d.Type+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio
	escuadrones -= 1
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	fmt.Println("Mensaje asincronico de laboratorio " + string(d.Body) + " leido.")

	defer connS.Close()
	var solicitudes = 0
	var mensaje_escuadron = ""
	serviceCliente := pb.NewMessageServiceClient(connS)
	flag := true
	if escuadrones == 1 {
		fmt.Println("Se envia escuadron 2 a Laboratorio " + string(d.Body))
		mensaje_escuadron = "2"
	} else {
		fmt.Println("Se envia escuadron 1 a Laboratorio " + string(d.Body))
		mensaje_escuadron = "1"
	}

	//Loop que se repite hasta que el laboratorio comunique a la central que todo esta OK
	for flag {
		//envia el mensaje al laboratorio
		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: mensaje_escuadron,
			})

		if err != nil {
			//panic("No se puede crear el mensaje " + err.Error())
			for i := 1; i <= 4; i++ {
				connS, err := grpc.Dial("dist04"+fmt.Sprint(i)+":50052", grpc.WithInsecure())
				if err != nil {
					//panic("No se pudo conectar con el servidor" + err.Error())
				} else {
					serviceCliente := pb.NewMessageServiceClient(connS)

					serviceCliente.Intercambio(context.Background(),
						&pb.Message{
							Body: "END",
						})
				}
			}
			os.Exit(1)
		}
		if res.Body == "Listo" {
			escuadrones += 1
			if escuadrones == 2 {
				fmt.Println("Retorno a Central Escuadra 2, Conexion Laboratorio " + string(d.Body) + " Cerrada.")
			} else {
				fmt.Println("Retorno a Central Escuadra 1, Conexion Laboratorio " + string(d.Body) + " Cerrada.")
			}
			flag = false
			ff.WriteString(string(d.Body) + " " + fmt.Sprint(solicitudes) + "\n")
			connS.Close()
		} else {
			solicitudes += 1
			if escuadrones == 1 {
				fmt.Println("Status Escuadra 2 : " + string(d.Body))
			} else {
				fmt.Println("Status Escuadra 1 : " + string(d.Body))
			}
		}
	}
}

func main() {
	qName := "Emergencias"                                           //Nombre de la cola
	hostQ := "localhost"                                             //Host de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	// ABRIENDO ARCHIVOS DE TEXTO
	file, err := os.Create("SOLICITUDES.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	//defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //Se crea la cola en RabbitMQ
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)

	//Ciclo que se repite infinatemente hasta la detencion del programa mediante ctrl+c, se esperan emergencias
	//y si se cuenta con al menos un escuadron se crea un Delivery para dar como input a la funcion 'todo' y poder
	//ejecutar esta ultima de manera concurrente
	for {
		fmt.Println("Esperando Emergencias")
		if escuadrones > 0 {
			chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
			if err != nil {
				log.Fatal(err)
			}
			for delivery := range chDelivery {
				go todo(delivery, file)
			}
		}
	}

}

// func init(): Cumple el proposito de cerrar los laboratorios en caso del termino del programa de la central mediante ctrl+c, enviando a cada una el mensaje de 'END' de esta forma cada Laboratorio realiza un os.Exit(1).
func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) //Se maneja el comando ctrl + c
	go func() {

		nombres := [4]string{"Laboratiorio Renca(la lleva) - Chile", "Laboratiorio Pohang - Korea", "Laboratiorio Kampala - Uganda", "Laboratiorio Pripiat - Rusia"}
		<-c
		for i := 1; i <= 4; i++ {
			connS, err := grpc.Dial("dist04"+fmt.Sprint(i)+":50052", grpc.WithInsecure())
			if err != nil {
				//panic("No se pudo conectar con el servidor" + err.Error())
			} else {
				serviceCliente := pb.NewMessageServiceClient(connS)

				serviceCliente.Intercambio(context.Background(),
					&pb.Message{
						Body: "END",
					})
				fmt.Println(nombres[i-1] + ": [APAGADO]")
			}
		}
		os.Exit(1)
	}()
}
