# SD - TP Sockets (Servidor de Broadcast Concurrente)

**Universidad Nacional de Misiones (UNaM)**  
**Facultad de Ciencias Exactas, Químicas y Naturales (FCEQyN)**  
**Asignatura:** Sistemas Distribuidos

**Profesores:**
* Ing. Rubén Luis María Castaño
* Lic. Claudio Omar Biale

---

## Grupo 1

* **Bauer, Luciano Agustín**
* **Olivieri, Ricardo**

---

## Ejecución Local

Para probar el proyecto en un entorno local, abrir terminales separadas en la raíz del proyecto y ejecutar:

```bash
# Terminal 1: servidor
docker-compose up --build

# Terminal 2: cliente (podés abrir varias terminales de cliente)
go run ./cmd/cliente/main.go
