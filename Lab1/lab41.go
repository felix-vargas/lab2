package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)


type server struct {
	pb.UnimplementedMessageServiceServer
}

//Función que calcula la probabilidad de que ocurra algo con i probabilidades a favor y devolverá falso si esa probabilidad se cumple. 
func probabilidad(i int) bool {
	rand.Seed(time.Now().UnixNano())
	var n int
	n = rand.Intn(10)
	if n < i {
		return false
	}
	return true
}

//Variables globales usadas para controlar las rutinas go de los labs
var resuelta = 0
var escuadron_llega = 0

//La función intercambio sirve para resolver los problemas cuando hay un estallido y también sirve para apagar el laboratorio.
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	if msg.Body == "END" {
		fmt.Println("APAGANDO LABORATORIO POR ORDEN DE LA CENTRAL")
		resuelta = 2
		os.Exit(1)
		return &pb.Message{Body: "Laboratiorio Renca(la lleva) - Chile: [APAGADO]"}, nil
	}

	if escuadron_llega == 0 {
		fmt.Println("Llega Escuadron " + msg.Body + ", conteniendo estallido...")
	}
	fmt.Println("Revisando estado Escuadron : ")
	escuadron_llega = 1
	for probabilidad(4) {
		fmt.Println("[NO LISTO]")
		time.Sleep(5 * time.Second)
		return &pb.Message{Body: "Laboratiorio Renca(la lleva) - Chile: [NO]"}, nil
	}
	resuelta = 1
	escuadron_llega = 0
	fmt.Println("[LISTO]")
	fmt.Println("Estallido contenido, Escuadron " + msg.Body + " Retornando.")
	return &pb.Message{Body: "Listo"}, nil
}

func main() {
	LabName := "Laboratiorio Renca(la lleva) - Chile"               //nombre del laboratorio
	qName := "Emergencias"                                          //nombre de la cola
	hostQ := "dist041"                                              //ip del servidor de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://test:test@" + hostQ + ":5672/") //conexion con RabbitMQ

	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	fmt.Println("Conexion a RabbitMQ exitosa")

	listener, err := net.Listen("tcp", ":50052") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMessageServiceServer(serv, &server{})
	go serv.Serve(listener)

	for {
    //Para apagar el laboratorio
		if resuelta == 2 {
			os.Exit(1)
		}
		fmt.Println("Analizando estado de Laboratorio: [ESTALLIDO/OK]")
		time.Sleep(5 * time.Second) //espera de 5 segundos
		//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)
		if probabilidad(2) {
			fmt.Println("SOS Enviado a Central. Esperando respuesta...")
			err = ch.Publish("", qName, false, false,
				amqp.Publishing{
					Headers:     nil,
					ContentType: "text/plain",
					Body:        []byte(LabName), //Contenido del mensaje
					Type:        "dist041",       // Entregar a la central que maquina envia el mensaje
				})
			resuelta = 0

			if err != nil {
				log.Fatal(err)
			}

      //Ciclo que nos servirá para esperar que se termine de ejecutar la función intercambio
			for resuelta == 0 {

			}
		}

	}

}
