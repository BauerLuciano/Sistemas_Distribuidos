package main

import (
	"fmt"
	"net"
	"os"
	
	"tp-sockets/internal/registro" 
)

func main() {
	reg := registro.NuevoRegistro()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Servidor concurrente escuchando en el puerto 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}

		go manejarConexion(conn, reg)
	}
}

func manejarConexion(conn net.Conn, reg *registro.Registro) {
	reg.Agregar(conn)
	
	defer func() {
		reg.Eliminar(conn)
		conn.Close()
	}()

	fmt.Printf("¡Nuevo cliente conectado desde: %s!\n", conn.RemoteAddr())
	
	fmt.Printf("Clientes activos en el registro: %d\n", len(reg.ObtenerTodos()))

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Cliente desconectado: %s\n", conn.RemoteAddr())
			return 
		}
	}
}