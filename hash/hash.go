package hash

import "hash/fnv"

const TAMAÑOINICIAL int = 31
const PORCENTAJEMAXIMO int = 50
const PORCENTAJEMINIMO int = 10
const FACTORREDIMENSION int = 3

type Hash struct {
	tamaño           int
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
	nodos := make([]*Nodo, TAMAÑOINICIAL)
	return &Hash{
		tamaño:           TAMAÑOINICIAL,
		cantidadOcupados: 0,
		cantidadBorrados: 0,
		nodos:            nodos,
	}
}

func (hash *Hash) Guardar(clave string, dato interface{}) {
	if (hash.cantidadBorrados+hash.cantidadOcupados)*100 > PORCENTAJEMAXIMO*hash.tamaño {
		hash.Redimensionar(hash.tamaño * FACTORREDIMENSION)
	}
	posicion := funcionHashing(clave) % uint32(hash.tamaño)
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
		if posicion >= uint32(hash.tamaño) {
			posicion = 0
		}
	}

	hash.nodos[posicion] = CrearNodo(clave, dato)
	hash.cantidadOcupados++

}

func (hash *Hash) Borrar(clave string) interface{} {
	if (hash.cantidadOcupados+hash.cantidadBorrados)*100 <= PORCENTAJEMINIMO*hash.tamaño &&
		hash.tamaño/FACTORREDIMENSION >= TAMAÑOINICIAL {
		hash.Redimensionar(hash.tamaño / FACTORREDIMENSION)
	}
	posicion := funcionHashing(clave) % uint32(hash.tamaño)
	for hash.nodos[posicion] != nil {
		if hash.nodos[posicion].clave == clave {
			if !hash.nodos[posicion].borrado {
				hash.nodos[posicion].borrado = true
				hash.cantidadBorrados++
				hash.cantidadOcupados--
				return hash.nodos[posicion].dato
			}
			break
		}
		posicion++
		if posicion >= uint32(hash.tamaño) {
			posicion = 0
		}
	}
	return nil
}

func (hash *Hash) Obtener(clave string) interface{} {

	posicion := funcionHashing(clave) % uint32(hash.tamaño)
	for hash.nodos[posicion] != nil {
		if hash.nodos[posicion].clave == clave {
			if hash.nodos[posicion].borrado {
				break
			}
			return hash.nodos[posicion].dato
		}
		posicion++
		if posicion >= uint32(hash.tamaño) {
			posicion = 0
		}
	}

	return nil

}

func (hash *Hash) Pertenece(clave string) bool {

	posicion := funcionHashing(clave) % uint32(hash.tamaño)
	for hash.nodos[posicion] != nil {
		if hash.nodos[posicion].clave == clave {
			if hash.nodos[posicion].borrado {
				break
			}
			return true
		}
		posicion++
		if posicion >= uint32(hash.tamaño) {
			posicion = 0
		}
	}

	return false
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

func (hash *Hash) Redimensionar(nuevoTamaño int) {
	nodosAnterior := hash.nodos
	nuevoSlice := make([]*Nodo, nuevoTamaño)
	hash.tamaño = nuevoTamaño
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
