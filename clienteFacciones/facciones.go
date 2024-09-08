package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Implementa la interfaz del servidor
type server struct{}

// Implementa el método del servicio
func (s *server) GetFaction(ctx context.Context, request *FactionRequest) (*FactionResponse, error) {
	// Lógica para obtener la facción solicitada
	faction := getFaction(request.Id)

	// Devuelve la respuesta
	return &FactionResponse{
		Id:          faction.Id,
		Name:        faction.Name,
		Description: faction.Description,
	}, nil
}

// Función para obtener una facción (solo para fines de demostración)
func getFaction(id int32) Faction {
	// Lógica para obtener la facción de una base de datos o cualquier otra fuente de datos
	// Aquí se devuelve una facción de ejemplo
	return Faction{
		Id:          id,
		Name:        "Ejemplo de facción",
		Description: "Esta es una facción de ejemplo",
	}
}

func main() {
	// Crea un servidor gRPC
	srv := grpc.NewServer()

	// Registra el servicio en el servidor
	RegisterFactionServiceServer(srv, &server{})

	// Escucha en el puerto 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	// Inicia el servidor
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
