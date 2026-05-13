package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("¡Conexión exitosa al servidor!")
	fmt.Println("Manteniendo la conexión abierta. (Presioná ENTER para salir)")

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}