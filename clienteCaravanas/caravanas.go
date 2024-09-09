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
		Las caravanas son las encargadas de realizar la entrega de los suministros en las llanuras de Eidolon.
		Como se mencion´o anteriormente, hay tres caravanas: dos dedicadas a los ostronitas y una destinada
		al transporte general de suministros Grineer y otros.

		Las caravanas permanecen en Cetus hasta que reciben alg´un paquete para entregar.

		Las caravanas esperan por un segundo paquete para entregar durante un periodo de tiempo
		definido al inicializar el sistema

		La caravana general, encargada de las entregas restantes, prioriza los paquetes prioritarios
		sobre los normales.

		Las caravanas siempre entregan primero el paquete que genere mayores ingresos para Konzu,
		a menos que el paquete no sea recibido. En tal caso, se intentar´a con el siguiente paquete y
		luego se intentar´a nuevamente entregar el paquete que no pudo ser entregado.

		Se considera una entrega completa cuando la caravana regresa a Cetus, incluso si no pudo
		entregar todos los paquetes.

		La caravana informa a Cetus cuando se completa una entrega, independientemente de si fue
		recibida o no

		Cada caravana lleva un registro de su funcionamiento, detallando cada paquete que tiene y ha
		entregado. Espec´ıficamente, ID del paquete, tipo de paquete, valor, escolta, destino, n´umero
		de intentos y fecha de entrega.

		Para las entregas realizadas por las caravanas existe un porcentaje de ´exito asociado a si se los
		suministro se logran entregar o se falla la misi´on, en este caso, existe un 85% de probabilidad de que
		se complete con ´exito la entrega de suministros al frente.

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
