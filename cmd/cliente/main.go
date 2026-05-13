package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"tp-sockets/pkg/protocolo"
)

func descubrirServidor() string {
	fmt.Println("📡 Buscando servidor en la red local...")
	serverAddr, _ := net.ResolveUDPAddr("udp", "255.255.255.255:9999")
	localAddr, _ := net.ResolveUDPAddr("udp", ":0")

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return "127.0.0.1:8080" 
	}
	defer conn.Close()

	conn.WriteToUDP([]byte("BUSCANDO_SERVER"), serverAddr)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	buf := make([]byte, 1024)
	_, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("⚠️  No se encontró servidor por UDP, probando localhost (Plan B)...")
		return "127.0.0.1:8080"
	}

	ipEncontrada := addr.IP.String()
	fmt.Printf("✅ ¡Servidor encontrado en la IP %s!\n", ipEncontrada)
	return fmt.Sprintf("%s:8080", ipEncontrada)
}

func main() {
	direccionServidor := descubrirServidor()

	conn, err := net.Dial("tcp", direccionServidor)
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("¡Conexión exitosa al servidor!")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingresá tu nombre de usuario: ")
	scanner.Scan()
	usuario := scanner.Text()

	fmt.Println("✅ ¡Todo listo! Empezá a chatear.")
	fmt.Print("> ")

	go func() {
		decoder := json.NewDecoder(conn)
		for {
			var msg protocolo.Mensaje
			if err := decoder.Decode(&msg); err != nil {
				fmt.Println("\n❌ El servidor se cerró o te desconectaste.")
				os.Exit(0)
			}
			fmt.Printf("\n[%s]: %s\n> ", msg.Usuario, msg.Texto)
		}
	}()

	encoder := json.NewEncoder(conn)
	for scanner.Scan() {
		texto := scanner.Text()
		if texto == "" {
			fmt.Print("> ")
			continue
		}

		msg := protocolo.Mensaje{
			Usuario: usuario,
			Texto:   texto,
		}

		if err := encoder.Encode(msg); err != nil {
			fmt.Println("Error enviando mensaje:", err)
			break
		}
		fmt.Print("> ")
	}
}