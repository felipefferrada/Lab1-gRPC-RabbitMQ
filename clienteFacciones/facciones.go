package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/felipefferrada/Lab1-gRPC-RabbitMQ/m/proto" // Asegúrate de cambiar esta línea con el path adecuado
	"google.golang.org/grpc"
)

func main() {
	// Establecer la conexión con el servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse al servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	// Enviar un mensaje al servidor

	/*
		Los clientes son quienes realizan las ordenes de suministros

		Generan ordenes periodicamente basadas en las necesidades de suministros de los Ostronitas o
		los Grineer, las cuales se envıan al sistema logıstico para su procesamiento. Para este proceso,
		se debe utilizar Protocol Buffers (ver gRPC).

		Los clientes pueden solicitar el estado de sus envıos ingresando y enviando su codigo de
		seguimiento al sistema logıstico, que les devolvera el estado de su pedido segun sus registros.
		Para este intercambio de informacion, se puede emplear cualquier medio de comunicacion entre
		el sistema logıstico y el cliente que sea confiable.

		Al momento de instanciarse cada sistema se requerir´a ingresar los tiempos bajo los cuales van
		a operar, ya sea el tiempo de espera de la caravana o cu´anto se demora en enviar cada paquete,
		y el tiempo de espera entre el env´ıo de ´ordenes en el cliente. Respecto a este ´ultimo, tambi´en
		debe solicitar el comportamiento a seguir (Ostronita/Grineer).

		Se debe generar una rutina para que simule el comportamiento de los clientes, considerando
		el env´ıo de ´ordenes y la solicitud del estado de un paquete.

	*/

	message := "Hola, servidor!"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SendMessage(ctx, &pb.MessageRequest{Message: message})
	if err != nil {
		log.Fatalf("Error al enviar el mensaje: %v", err)
	}

	log.Printf("Respuesta del servidor: %s", resp.GetReply())
}
