package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)





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



func DateNodeRandom() (Nombre_DateNode string, IP string) {
	rand.Seed(time.Now().UnixNano())
	switch os := rand.Intn(3); os {
	case 0:
		Nombre := "DateNode Grunt"
		IP := "dist042:50051"
		return Nombre, IP
	case 1:
		Nombre := "DateNode Synth"
		IP := "dist043:50051"
		return Nombre, IP
	default:
		Nombre := "DateNode Cremator"
		IP := "dist044:50051"
		return Nombre, IP
	}
}








func GuardarDATA(data string) {

	Split_Msj := strings.Split(data, ":")
	Tipo := Split_Msj[0]
	ID := Split_Msj[1]

	NAMEDATENODE, IPNODE := DateNodeRandom()

	_, err := file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	err = file.Sync()

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	connS, err := grpc.Dial(IPNODE, grpc.WithInsecure())
    
    
    if err != nil {
        panic("No se pudo conectar con el servidor" + err.Error())
    }

    
    defer connS.Close()
    
    serviceCliente := pb.NewMessageServiceClient(connS)


    //envia el mensaje al laboratorio
    res, err := serviceCliente.Intercambio(context.Background(), 
        &pb.Message{
            Body: data,
            Type: false,
        })

    if err != nil {
        panic("No se puede crear el mensaje " + err.Error())
    }
    
    fmt.Println(res.Body) //respuesta del laboratorio
    time.Sleep(5 * time.Second) //espera de 5 segundos

}

func Fetch_Rebeldes( tipo string) string{

    var Respuesta []string

    
    //CONEXION DATANODE 1 
    connS, err := grpc.Dial("dist042:50051", grpc.WithInsecure())
    if err != nil {
        panic("No se pudo conectar con el servidor" + err.Error())
    }
    defer connS.Close()

    serviceCliente := pb.NewMessageServiceClient(connS)
    //envia el mensaje al laboratorio
    res, err := serviceCliente.Intercambio(context.Background(), 
        &pb.Message{
            Body: tipo,
            Type: true,
        })
    if err != nil {
        panic("No se puede crear el mensaje " + err.Error())
    }
    append(Respuesta,res.Body)


    
    //CONEXION DATANODE 2
        connS, err := grpc.Dial("dist043:50051", grpc.WithInsecure())
    if err != nil {
        panic("No se pudo conectar con el servidor" + err.Error())
    }

    defer connS.Close()

    serviceCliente := pb.NewMessageServiceClient(connS)
    //envia el mensaje al laboratorio
    res, err := serviceCliente.Intercambio(context.Background(), 
        &pb.Message{
            Body: tipo,
            Type: true,
        })
    if err != nil {
        panic("No se puede crear el mensaje " + err.Error())
    }
    append(Respuesta,res.Body)


    
    //CONEXION DATANODE 3
            connS, err := grpc.Dial("dist044:50051", grpc.WithInsecure())
    if err != nil {
        panic("No se pudo conectar con el servidor" + err.Error())
    }

    defer connS.Close()
    serviceCliente := pb.NewMessageServiceClient(connS)
    //envia el mensaje al laboratorio
    res, err := serviceCliente.Intercambio(context.Background(), 
        &pb.Message{
            Body: tipo,
            Type: true,
        })
    if err != nil {
        panic("No se puede crear el mensaje " + err.Error())
    }
    append(Respuesta,res.Body)

    return Respuesta
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
    if (msg.Type == true){//conbine
        Split_Msj := strings.Split(msg.Body,":")
        ID := Split_Msj[1]
        if (RevisarID(ID)){         
            GuardarDATA(msg.Body)
            msn := "Guardado"   
        }else{
            msn := "ID Repetido" 
        }
    }else{//rebelde
        msn :=Fetch_Rebeldes(msg.Body)
    }
    return &pb.Message{Body: msn}, nil 
}




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
