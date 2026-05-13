package registro

import (
	"net"
	"sync"
)

type Registro struct {
	clientes map[net.Conn]bool
	mu       sync.RWMutex 
}

// NuevoRegistro inicializa la estructura
func NuevoRegistro() *Registro {
	return &Registro{
		clientes: make(map[net.Conn]bool),
	}
}

// Agregar registra una nueva conexión
func (r *Registro) Agregar(conn net.Conn) {
	r.mu.Lock()         
	defer r.mu.Unlock() 
	
	r.clientes[conn] = true
}

// Eliminar quita una conexión cuando un cliente se va
func (r *Registro) Eliminar(conn net.Conn) {
	r.mu.Lock()  
	defer r.mu.Unlock()
	
	delete(r.clientes, conn)
}

func (r *Registro) ObtenerTodos() []net.Conn {
	r.mu.RLock()        
	defer r.mu.RUnlock()

	var activos []net.Conn
	for conn := range r.clientes {
		activos = append(activos, conn)
	}
	return activos
}