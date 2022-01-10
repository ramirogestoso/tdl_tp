package hash

import "hash/fnv"

const tamanoINICIAL int = 31
const PORCENTAJEMAXIMO int = 50
const PORCENTAJEMINIMO int = 10
const FACTORREDIMENSION int = 3

type Hash struct {
	tamano           int
	cantidadOcupados int
	cantidadBorrados int
	nodos            []*Nodo
}

type Nodo struct {
	clave   string
	dato    interface{}
	borrado bool
}

func CrearNodo(clave string, dato interface{}) *Nodo {
	return &Nodo{
		clave:   clave,
		dato:    dato,
		borrado: false,
	}
}

func CrearHash() *Hash {
	nodos := make([]*Nodo, tamanoINICIAL)
	return &Hash{
		tamano:           tamanoINICIAL,
		cantidadOcupados: 0,
		cantidadBorrados: 0,
		nodos:            nodos,
	}
}

func (hash *Hash) Guardar(clave string, dato interface{}) {
	if (hash.cantidadBorrados+hash.cantidadOcupados)*100 > PORCENTAJEMAXIMO*hash.tamano {
		hash.redimensionar(hash.tamano * FACTORREDIMENSION)
	}
	posicion := funcionHashing(clave) % uint32(hash.tamano)
	for hash.nodos[posicion] != nil {
		if hash.nodos[posicion].clave == clave {
			hash.nodos[posicion].dato = dato
			if hash.nodos[posicion].borrado {
				hash.nodos[posicion].borrado = false
				hash.cantidadBorrados--
				hash.cantidadOcupados++
			}
			return
		}
		posicion++
		if posicion >= uint32(hash.tamano) {
			posicion = 0
		}
	}

	hash.nodos[posicion] = CrearNodo(clave, dato)
	hash.cantidadOcupados++

}

func (hash *Hash) Borrar(clave string) interface{} {
	if (hash.cantidadOcupados+hash.cantidadBorrados)*100 <= PORCENTAJEMINIMO*hash.tamano &&
		hash.tamano/FACTORREDIMENSION >= tamanoINICIAL {
		hash.redimensionar(hash.tamano / FACTORREDIMENSION)
	}

	nodo := hash.obtenerNodo(clave)
	if nodo == nil { return nil } // no existe o esta borrado ya
	nodo.borrado = true
	hash.cantidadBorrados++
	hash.cantidadOcupados--
	return nodo.dato
}

func (hash *Hash) Obtener(clave string) interface{} {
	nodo := hash.obtenerNodo(clave)
	if nodo == nil { return nil } // no existe o esta borrado
	return nodo.dato
}

func (hash *Hash) Pertenece(clave string) bool {
	nodo := hash.obtenerNodo(clave)
	if nodo == nil { return false } // no existe o esta borrado
	return true
}

func (hash *Hash) obtenerNodo(clave string) *Nodo {
	posicion := funcionHashing(clave) % uint32(hash.tamano)
	for hash.nodos[posicion] != nil {
		if hash.nodos[posicion].clave == clave {
			if hash.nodos[posicion].borrado {
				break
			}
			return hash.nodos[posicion]
		}
		posicion++
		if posicion >= uint32(hash.tamano) {
			posicion = 0
		}
	}
	return nil
}

func (hash *Hash) DatoPertenece(dato interface{}) bool {
	for _, nodo := range hash.nodos {
		if nodo != nil && nodo.dato == dato && !nodo.borrado {
			return true
		}
	}
	return false
}

func (hash *Hash) Largo() int {
	return hash.cantidadOcupados
}

func (hash *Hash) redimensionar(nuevotamano int) {
	nodosAnterior := hash.nodos
	nuevoSlice := make([]*Nodo, nuevotamano)
	hash.tamano = nuevotamano
	hash.cantidadBorrados = 0
	hash.cantidadOcupados = 0
	hash.nodos = nuevoSlice
	for _, nodo := range nodosAnterior {
		if nodo != nil && !nodo.borrado {
			hash.Guardar(nodo.clave, nodo.dato)
		}
	}
}

func funcionHashing(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
