package main

import (
	"context"
	"fmt"
    "strings"
    "strconv"
    
    pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)



func SolicitarInput() []string {
    for{
        fmt.Println("> POR FAVOR ESCRIBIR SOLICITUD : ")
        var solicitud string
        fmt.Scanln(&solicitud)
        var solicitud_lista = strings.Split(solicitud,":")

        //Verificar formato de input sea correcto, primero largo
        if(len(solicitud_lista) == 3){
            //Luego si los tipos calzan
            if(solicitud_lista[0] == "LOGISTICA" || solicitud_lista[0] == "FINANCIERA" || solicitud_lista[0] == "MILITAR"){
                //Si el ID es un numero
                if _, err := strconv.Atoi(solicitud_lista[1]); err == nil {
                    return solicitud_lista
                }
            }
        }

    }
}




func main(){

    connS, err := grpc.Dial("dist041:50052", grpc.WithInsecure())
    serviceCliente := pb.NewMessageServiceClient(connS)

    for{
        var input []string
    
        input = SolicitarInput()
    
        if err != nil{
            res, _ := serviceCliente.Intercambio(context.Background(),
        			&pb.Message{
        				Body: input[0]+":"+input[1]+":"+input[2],
        	        })
            if res == "SI"{
                fmt.Println("ENTRADA CREADA EXITOSAMENTE")
            }
            if res == "NO"{
                fmt.Println("NO SE HA PODIDO CREAR LA ENTRADA > ID REPETIDO ")
            }
        }
    }


    
}