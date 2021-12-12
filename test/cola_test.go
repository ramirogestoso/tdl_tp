package test

import (
  "testing"
  "github.com/ramirogestoso/tp/cola"
)

func TestColaVacia(t *testing.T) {
  c := cola.CrearCola()
  ok := c.EstaVacia()
  if !ok {t.Error()}
}

func TestColaConUnElementoNoEstaVacia(t *testing.T) {
  c := cola.CrearCola()
  c.Encolar(1)
  ok := !c.EstaVacia()
  if !ok {t.Error()}
}

func TestColaConUnElementoTopeEstaBien(t *testing.T) {
  c := cola.CrearCola()
  c.Encolar(1)
  ok := c.VerTope() == 1
  if !ok {t.Error()}
}

func TestColaConVariosElementosElTopeSeMantiene(t *testing.T) {
  c := cola.CrearCola()
  ok := true
  for i:=0; i<10; i++ {
    c.Encolar(i)
    ok = ok && c.VerTope() == 0
  }
  if !ok {t.Error()}
}

func TestVerTopeEnColaVaciaEsNil(t *testing.T) {
  c := cola.CrearCola()
  ok := c.VerTope() == nil
  if !ok {t.Error()}
}


func TestDesencolarEnColaVaciaEsNil(t *testing.T) {
  c := cola.CrearCola()
  ok := c.Desencolar() == nil
  if !ok {t.Error()}
}

func TestDesencolarConUnElemento(t *testing.T) {
  c := cola.CrearCola()
  c.Encolar(12)
  ok := c.Desencolar() == 12
  if !ok {t.Error()}
}

func TestEncolarDesencolarEstaVacia(t *testing.T) {
  c := cola.CrearCola()
  c.Encolar(nil)
  c.Desencolar()
  ok := c.EstaVacia()
  if !ok {t.Error()}
}

func TestEncolarDesencolarHastaVaciarEstaVacia(t *testing.T) {
  c := cola.CrearCola()
  for i:=0; i<10; i++ {
    c.Encolar(i)
  }
  for i:=0; i<10; i++ {
    c.Desencolar()
  }
  ok := c.EstaVacia()
  if !ok {t.Error()}
}

func TestEncolarYDesencolarDevuelveCorrectamente(t *testing.T) {
  c := cola.CrearCola()
  ok := true
  for i:=0; i<10; i++ {
    c.Encolar(i)
  }
  for i:=0; i<10; i++ {
    ok = ok && c.Desencolar() == i
  }
  if !ok {t.Error()}
}

func TestEncolarNilEsValido(t *testing.T) {
  c := cola.CrearCola()
  c.Encolar(nil)
  ok := !c.EstaVacia() && c.VerTope() == nil && c.Desencolar() == nil
  if !ok {t.Error()}
}
