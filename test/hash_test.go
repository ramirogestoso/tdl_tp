package test

import (
	"github.com/ramirogestoso/tp/hash"
	"testing"
)

func TestHashVacio(t *testing.T) {
	hash := hash.CrearHash()
	ok := hash.Largo() == 0

	if !ok {
		t.Error()
	}
}

func TestNuevaClavePertenece(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	ok := hash.Pertenece("mascota")

	if !ok {
		t.Error()
	}
}

func TestDatoPertenece(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	hash.Guardar("edad", "24")
	hash.Guardar("genero", "Masculino")
	ok := hash.DatoPertenece("Perro") && hash.DatoPertenece("24") && hash.DatoPertenece("Masculino")

	if !ok {
		t.Error()
	}
}

func TestDatoNoPertenece(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	hash.Guardar("edad", "24")
	hash.Guardar("genero", "Masculino")
	ok := !hash.DatoPertenece("Gato")

	if !ok {
		t.Error()
	}
}

func TestNuevaClaveObtener(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	ok := hash.Obtener("mascota") == "Perro"

	if !ok {
		t.Error()
	}
}

func TestTresNuevasClavesObtener(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	hash.Guardar("edad", "24")
	hash.Guardar("genero", "Masculino")
	ok := hash.Obtener("mascota") == "Perro" &&
		hash.Obtener("edad") == "24" &&
		hash.Obtener("genero") == "Masculino"

	if !ok {
		t.Error()
	}
}

func TestNuevasClavesTamañoHash(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	hash.Guardar("edadd", "24")
	hash.Guardar("genero", "Masculino")

	ok := hash.Largo() == 3

	if !ok {
		t.Error()
	}
}

func TestBorrarClaveTamañoHash(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	hash.Guardar("edad", "24")
	hash.Guardar("genero", "Masculino")
	hash.Borrar("mascota")

	ok := hash.Largo() == 2

	if !ok {
		t.Error()
	}
}

func TestClaveInexistente(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascotta", "Perro")
	hash.Guardar("cualquierEdad", "24")
	hash.Guardar("genero", "Masculino")
	ok := !hash.Pertenece("nacionalidad")

	if !ok {
		t.Error()
	}
}

func TestBorrarClave(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascota", "Perro")
	hash.Guardar("edad", "24")
	hash.Guardar("generro", "Masculino")
	hash.Borrar("mascota")
	ok := !hash.Pertenece("mascota")

	if !ok {
		t.Error()
	}
}

func TestBorrarClaveNoBorraOtras(t *testing.T) {
	hash := hash.CrearHash()
	hash.Guardar("mascotaa", "Perro")
	hash.Guardar("edad", "24")
	hash.Guardar("genero", "Masculino")
	hash.Borrar("mascotaa")
	ok := hash.Pertenece("edad") && hash.Pertenece("genero") && !hash.Pertenece("mascotaa")

	if !ok {
		t.Error()
	}
}

func TestRedimensionMuchosElementos(t *testing.T) {
	hash := hash.CrearHash()
	clave := "mascota"
	for i := 0; i < 35; i++ {
		clave = clave + "1"
		hash.Guardar(clave, "Perro")
	}

	ok := hash.Largo() == 35

	if !ok {
		t.Error()
	}
}

func TestRedimensionConBorrado(t *testing.T) {
	hash := hash.CrearHash()
	clave := "mascota"
	for i := 0; i < 5; i++ {
		clave = clave + "1"
		hash.Guardar(clave, "Perro")
	}

	hash.Borrar("mascota1")

	for i := 5; i < 35; i++ {
		clave = clave + "1"
		hash.Guardar(clave, "Perro")
	}

	ok := !hash.Pertenece("mascota1") && hash.Largo() == 34

	if !ok {
		t.Error()
	}
}

func TestRedimensionNegativa(t *testing.T) {
	hash := hash.CrearHash()
	clave := "mascota"
	for i := 0; i < 16; i++ {
		clave = clave + "1"
		hash.Guardar(clave, "Perro")
	}

	clave = "mascota"
	for i := 0; i < 7; i++ {
		clave = clave + "1"
		hash.Borrar(clave)
	}

	ok := hash.Largo() == 9

	if !ok {
		t.Error()
	}
}

func TestRedimensionNegativaGrande(t *testing.T) {
	hash := hash.CrearHash()
	clave := "mascota"
	for i := 0; i < 47; i++ {
		clave = clave + "1"
		hash.Guardar(clave, "Perro")
	}

	cond1 := hash.Largo() == 47

	clave = "mascota"
	for i := 0; i < 23; i++ {
		clave = clave + "1"
		hash.Borrar(clave)
	}

	cond2 := hash.Largo() == 24

	ok := cond1 && cond2

	if !ok {
		t.Error()
	}
}
