package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
    "net"
    "fmt"
    "context"

    
    pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)


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




type server struct {
	pb.UnimplementedMessageServiceServer
}



func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
    file, err := os.Create("DATA.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
    if (msg.Type == true){
        msn := RetornarData (msg.Body)
    }else{
        file.WriteString(msg.Body + "\n")
        msn := "Guardado"
    }
    
    return &pb.Message{Body: msn}, nil
    
}



//"DateNode Cremator"
func main() {



    listener, err := net.Listen("tcp", ":50051") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	for {
		pb.RegisterMessageServiceServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}
    
}