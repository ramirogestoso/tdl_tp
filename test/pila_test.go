package pila_test

import (
  "github.com/ramirogestoso/tp/pila"
  "testing"
)

func TestPilaVacia(t *testing.T) {
  p := pila.CrearPila()
  ok := p.EstaVacia()
  if !ok {
    t.Error()
  }
}

func TestPilaUnElementoNoEstaVacia(t *testing.T) {
  p := pila.CrearPila()
  p.Apilar(1)
  ok := !p.EstaVacia()
  if !ok {t.Error()}
}

func TestPilaDosElementosNoEstaVacia(t *testing.T) {
  p := pila.CrearPila()
  p.Apilar("uno")
  p.Apilar(1)
  p.Desapilar()
  ok := !p.EstaVacia()
  if !ok {t.Error()}
}

func TestPilaDesapilaEnOrdenLIFO(t *testing.T) {
  p := pila.CrearPila()
  p.Apilar("uno")
  p.Apilar(1)
  ok := p.Desapilar() == 1 && p.Desapilar() == "uno"
  if !ok {t.Error()}
}

func TestPilaApilarDesapilarEstaVacia(t *testing.T) {
  p := pila.CrearPila()
  p.Apilar("uno")
  p.Desapilar()
  ok := p.EstaVacia()
  if !ok {t.Error()}
}

func TestPilaVerTopeDevuelveBien(t *testing.T) {
  p := pila.CrearPila()
  var x int = 12
  p.Apilar(&x)
  ok := p.VerTope() == &x
  if !ok {t.Error()}
}

func TestPilaApilarNilNoEstaVacia(t *testing.T) {
  p := pila.CrearPila()
  p.Apilar(nil)
  ok := !p.EstaVacia()
  if !ok {t.Error()}
}

func TestPilaDesapilarPilaVaciaDevuelveNil(t *testing.T) {
  p := pila.CrearPila()
  ok := p.Desapilar() == nil
  if !ok {t.Error()}
}

func TestVerTopeDePilaVaciaEsNil(t *testing.T) {
  p := pila.CrearPila()
  ok := p.VerTope() == nil
  if !ok {t.Error()}
}
