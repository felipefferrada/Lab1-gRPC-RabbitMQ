package main

import (
	"context"
	"log"
	"net"

	pb "github.com/felipefferrada/Lab1-gRPC-RabbitMQ/m/proto" // Asegúrate de cambiar esta línea con el path adecuado

	"google.golang.org/grpc"
)

// Implementar el servidor
type server struct {
	pb.UnimplementedMessageServiceServer
}

// Implementar la función SendMessage
func (s *server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Mensaje recibido: %s", req.GetMessage())
	// Responder al cliente
	return &pb.MessageResponse{Reply: "Mensaje recibido correctamente"}, nil
}

func main() {
	// Escuchar en un puerto específico
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	// Crear un servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar el servicio MessageService
	pb.RegisterMessageServiceServer(grpcServer, &server{})

	log.Printf("Servidor escuchando en el puerto 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
