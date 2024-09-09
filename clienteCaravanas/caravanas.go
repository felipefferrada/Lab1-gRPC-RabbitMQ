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
