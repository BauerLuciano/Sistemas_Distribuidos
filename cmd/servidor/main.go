package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"tp-sockets/internal/registro"
	"tp-sockets/pkg/protocolo"
)


func escucharUDP() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error iniciando oreja UDP:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		msg := string(buf[:n])
		if msg == "BUSCANDO_SERVER" {
			fmt.Printf("¡Grito UDP recibido de %s! Respondiendo...\n", clientAddr.IP.String())
			conn.WriteToUDP([]byte("ACA_ESTOY"), clientAddr)
		}
	}
}

func main() {
	reg := registro.NuevoRegistro()

	go escucharUDP()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("🚀 Servidor concurrente escuchando en el puerto 8080...")

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
		fmt.Printf("Clientes activos en el registro: %d\n", len(reg.ObtenerTodos()))
	}()

	fmt.Printf("¡Nuevo cliente conectado desde: %s!\n", conn.RemoteAddr())
	fmt.Printf("Clientes activos en el registro: %d\n", len(reg.ObtenerTodos()))
	
	decoder := json.NewDecoder(conn)

	for {
		var msg protocolo.Mensaje
		if err := decoder.Decode(&msg); err != nil {
			fmt.Printf("Cliente desconectado: %s\n", conn.RemoteAddr())
			return
		}

		for _, cliente := range reg.ObtenerTodos() {
			if cliente != conn {
				encoder := json.NewEncoder(cliente)
				encoder.Encode(msg)
			}
		}
	}
}