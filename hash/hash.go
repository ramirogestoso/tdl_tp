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
	
	pos := hash.obtenerPosicionNodo(clave)
	nodo := hash.nodos[pos]
	if nodo != nil /* la clave ya existe */ {
		hash.nodos[pos].dato = dato /* actualiza el dato que tenia */
		if nodo.borrado {
			//hash.nodos[pos].borrado = false /* deshace borrar */
			nodo.borrado = false
			hash.cantidadBorrados--
		}
	} else { /* la clave no existe y ya tenemos la posicion deseada */
		hash.nodos[pos] = CrearNodo(clave, dato)
	}
	hash.cantidadOcupados++
}

func (hash *Hash) Borrar(clave string) interface{} {
	if (hash.cantidadOcupados+hash.cantidadBorrados)*100 <= PORCENTAJEMINIMO*hash.tamano &&
		hash.tamano/FACTORREDIMENSION >= tamanoINICIAL {
		hash.redimensionar(hash.tamano / FACTORREDIMENSION)
	}

	pos := hash.obtenerPosicionNodo(clave)
	nodo := hash.nodos[pos]
	if nodo == nil || nodo.borrado { return nil }
	hash.nodos[pos].borrado = true
	hash.cantidadBorrados++
	hash.cantidadOcupados--
	return hash.nodos[pos].dato
}

func (hash *Hash) Obtener(clave string) interface{} {
	pos := hash.obtenerPosicionNodo(clave)
	nodo := hash.nodos[pos]
	if nodo == nil || nodo.borrado { return nil }
	return hash.nodos[pos].dato
}

func (hash *Hash) Pertenece(clave string) bool {
	pos := hash.obtenerPosicionNodo(clave)
	nodo := hash.nodos[pos]
	if nodo == nil || nodo.borrado { return false }
	return true
}

func (hash *Hash) obtenerPosicionNodo(clave string) int {
	/* devuelve la posicion donde deberia estar el nodo */
	posicion := funcionHashing(clave) % uint32(hash.tamano)
	for hash.nodos[posicion] != nil && hash.nodos[posicion].clave != clave {
		posicion++
		if posicion >= uint32(hash.tamano) {
			posicion = 0
		}
	}
	return int(posicion)
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
