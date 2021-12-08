package test

import (
  "testing"
  "../lista"
)

func TestListaNuevaEstaVacia(t *testing.T) {
  l := lista.CrearLista()
  ok := l.EstaVacia()
  if !ok {t.Error()}
}

func TestListaConUnElementoNoEstaVacia(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  ok := !l.EstaVacia()
  if !ok {t.Error()}
}

func TestListaConUnElementoObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  x := l.ObtenerEnPosicion(0)
  ok := x == 3
  if !ok {t.Error()}
}

func TestListaDosElementosObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  l.InsertarUltimo(7)
  x := l.ObtenerEnPosicion(0)
  ok := x == 3
  if !ok {t.Error()}
}

func TestListaTresElementosObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  l.InsertarUltimo(7)
  l.InsertarUltimo(10)
  x := l.ObtenerEnPosicion(2)
  ok := x == 10
  if !ok {t.Error()}
}

func TestInsertarPrimeroObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  l.InsertarUltimo(7)
  l.InsertarPrimero(10)
  x := l.ObtenerEnPosicion(0)
  ok := x == 10
  if !ok {t.Error()}
}

func TestInsertarVariosAlInicioObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarPrimero(3)
  l.InsertarPrimero(7)
  l.InsertarPrimero(10)
  x := l.ObtenerEnPosicion(2)
  ok := x == 3
  if !ok {t.Error()}
}

func TestListaInsertarEnPosicionObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  l.InsertarUltimo(7)
  l.InsertarUltimo(10)
  l.InsertarEnPosicion(1, 2)
  x := l.ObtenerEnPosicion(1)
  ok := x == 2
  if !ok {t.Error()}
}

func TestListaInsertarEnPosicionFinalObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  l.InsertarUltimo(7)
  l.InsertarUltimo(10)
  l.InsertarEnPosicion(l.ObtenerLargo(), 2)
  x := l.ObtenerEnPosicion(l.ObtenerLargo()-1)
  ok := x == 2
  if !ok {t.Error()}
}

func TestListaInsertarEnPosicionInicialObtenerElementoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(3)
  l.InsertarUltimo(7)
  l.InsertarUltimo(10)
  l.InsertarEnPosicion(0, 2)
  x := l.ObtenerEnPosicion(0)
  ok := x == 2
  if !ok {t.Error()}
}

func TestListaConDosElementosTieneLargoCorrecto(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(1)
  l.InsertarUltimo(2)
  ok := l.ObtenerLargo() == 2
  if !ok {t.Error()}
}

func TestEliminarElementoDeListaConVariosElementosTieneMenorLargo(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  l.InsertarUltimo("hola")
  largo1 := l.ObtenerLargo()
  l.RemoverEnPosicion(1)
  ok := l.ObtenerLargo() < largo1
  if !ok {t.Error()}
}

func TestEliminarElementoDeListaConVariosElementosDevuelveCorrectamente(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo("hola")
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  removido := l.RemoverEnPosicion(1)
  ok := removido == "hola"
  if !ok {t.Error()}
}

func TestEliminarPrimerElementoDevuelveCorrectamente(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo("hola")
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  removido := l.RemoverEnPosicion(1)
  ok := removido == "hola"
  if !ok {t.Error()}
}

func TestEliminarElementoNoEstaEnLista(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo("hola")
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  removido := l.RemoverEnPosicion(1)
  ok := removido != l.ObtenerEnPosicion(0)
  if !ok {t.Error()}
}

func TestEliminarElementoDelMedioFuncionaCorrectamente(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo("hola")
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  l.RemoverEnPosicion(1)
  ok := l.ObtenerEnPosicion(0) == nil && l.ObtenerEnPosicion(1) == 2
  if !ok {t.Error()}
}

func TestEliminarUltimoElementoFuncionaCorrectamente(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo("hola")
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  l.RemoverUltimo()
  ok := l.ObtenerUltimo() == 2
  if !ok {t.Error()}
}

func TestEliminarPrimeroElementoFuncionaCorrectamente(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(nil)
  l.InsertarUltimo("hola")
  l.InsertarUltimo(2)
  l.InsertarUltimo(15)
  l.RemoverPrimero()
  ok := l.ObtenerEnPosicion(0) == "hola"
  if !ok {t.Error()}
}

func TestEliminarUltimoDeListaVacia(t *testing.T) {
  l := lista.CrearLista()
  ok := l.RemoverUltimo() == nil
  if !ok {t.Error()}
}

func TestEliminarPrimeroDeListaVacia(t *testing.T) {
  l := lista.CrearLista()
  ok := l.RemoverPrimero() == nil
  if !ok {t.Error()}
}

func TestObtenerConIndiceNegativoEsNil(t *testing.T) {
  l := lista.CrearLista()
  ok := l.ObtenerEnPosicion(-2) == nil
  if !ok {t.Error()}
}

func TestEncontrarElemento(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(1)
  l.InsertarUltimo(2)
  l.InsertarUltimo(3)
  ok := l.Encontrar(2) == 1
  if !ok {t.Error()}
}

func TestEncontrarElementoQueNoExiste(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(1)
  l.InsertarUltimo(2)
  l.InsertarUltimo(3)
  ok := l.Encontrar("char") == -1
  if !ok {t.Error()}
}

func TestEncontrarElementoEnListaVacia(t *testing.T) {
  l := lista.CrearLista()
  ok := l.Encontrar(2) == -1
  if !ok {t.Error()}
}

/// test iter

func TestIteradorListaVaciaNoCambiaPosicion(t *testing.T) {
  l := lista.CrearLista()
  it := l.CrearIterador()
  ok := !it.Avanzar() && !it.Retroceder()
  if !ok {t.Error()}
}

func TestIteradorListaConUnElemento(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(0)
  it := l.CrearIterador()
  ok := it.VerActual() == 0
  if !ok {t.Error()}
}


func TestIteradorAvanzarAlFinal(t *testing.T) {
  l := lista.CrearLista()
  l.InsertarUltimo(0)
  l.InsertarUltimo(4)
  it := l.CrearIterador()
  it.Avanzar()
  it.Avanzar()
  ok := !it.Avanzar() && it.VerActual() == nil
  if !ok {t.Error()}
}

func TestListaLargaIteradorAvanzarAlFinal(t *testing.T) {
  l := lista.CrearLista()
  for i:=0; i<1000; i++ {
    l.InsertarUltimo(i)
  }
  it := l.CrearIterador()
  for it.Avanzar() { }
  ok := it.VerActual() == nil
  if !ok {t.Error()}
}

func TestListaLargaIteradorVerUltimo(t *testing.T) {
  l := lista.CrearLista()
  for i:=0; i<1000; i++ {
    l.InsertarUltimo(i)
  }
  it := l.CrearIterador()
  for it.Avanzar() { }
  it.Retroceder()
  ok := it.VerActual() == 999
  if !ok {t.Error()}
}

func TestListaLargaIteradorAvanzarAlFinalYVolverAlInicio(t *testing.T) {
  l := lista.CrearLista()
  for i:=0; i<1000; i++ {
    l.InsertarUltimo(i)
  }
  it := l.CrearIterador()
  for it.Avanzar() {}
  for it.Retroceder() {}
  ok := it.VerActual() == nil && it.Avanzar() && it.VerActual() == 0
  if !ok {t.Error()}
}

func TestListaLargaIteradorVerActuales(t *testing.T) {
  l := lista.CrearLista()
  for i:=0; i<1000; i++ {
    l.InsertarUltimo(i)
  }
  it := l.CrearIterador()
  k := 0
  ok := it.VerActual() == k
  for it.Avanzar() {
    k++
    ok = ok && it.VerActual() == k
  }
  if !ok {t.Error()}
}

func TestIteradorAvanzarYRetroceder(t *testing.T) {
  l := lista.CrearLista()
  for i:=0; i<5; i++ {
    l.InsertarUltimo(i)
  }
  it := l.CrearIterador()

  ok := it.VerActual() == 0
  it.Avanzar()
  ok = ok && it.VerActual() == 1
  it.Avanzar()
  ok = ok && it.VerActual() == 2
  it.Avanzar()
  ok = ok && it.VerActual() == 3
  it.Retroceder()
  ok = ok && it.VerActual() == 2
  it.Retroceder()
  ok = ok && it.VerActual() == 1
  it.Avanzar()
  ok = ok && it.VerActual() == 2
  it.Retroceder()
  it.Retroceder()
  ok = ok && it.VerActual() == 0
  ok = ok && !it.Retroceder()
  ok = ok && it.VerActual() == nil

  if !ok {t.Error()}
}
