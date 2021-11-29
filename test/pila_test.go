package pila_test

import (
  "../pila"
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

func TestPilaDesapilarPilaVaciaPanic(t *testing.T) {
  p := pila.CrearPila()
  defer func() {
    if r:=recover(); r == nil {t.Error()}
  }()
  p.Desapilar()
}
