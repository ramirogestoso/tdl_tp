package test

import (
  "testing"
  "github.com/ramirogestoso/tp/abb"
)

func cmpString(c1 interface{}, c2 interface{}) int {
  clave1 := c1.(string)
  clave2 := c2.(string)
  switch {
  case clave2>clave1: return -1
  case clave1>clave2: return 1
  default: return 0
  }
}

func cmpInt(c1 interface{}, c2 interface{}) int {
  return c1.(int) - c2.(int)
}

func TestAbbNuevoTieneCantidadCero(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  ok := a.Cantidad() == 0
  if !ok {t.Error()}
}

func TestAbbInsertarUnElementoCantidadEsUno(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("uno", 1)
  ok := a.Cantidad() == 1
  if !ok {t.Error()}
}

func TestAbbInsertarVariosCantidadCorrecta(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("g", nil)
  a.Insertar("v", nil)
  a.Insertar("c", nil)
  a.Insertar("f", nil)
  a.Insertar("m", nil)
  a.Insertar("s", nil)
  a.Insertar("a", nil)
  a.Insertar("b", nil)
  ok := a.Cantidad() == 8
  if !ok {t.Error()}
}

func TestAbbConUnElementoPertenece(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("Hola Mundo", 12)
  ok := a.Pertenece("Hola Mundo")
  if !ok {t.Error()}
}

func TestAbbConVariosElementosPertenecen(t *testing.T) {
  a := abb.CrearAbb(cmpInt)
  a.Insertar(1, 1)
  a.Insertar(2, 2)
  a.Insertar(3, 3)
  ok := a.Pertenece(1) && a.Pertenece(2) && a.Pertenece(3)
  if !ok {t.Error()}
}

func TestAbbConVariosElementosNoPertenece(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("uno", 1)
  a.Insertar("dos", 2)
  a.Insertar("tres", 3)
  ok := !a.Pertenece("unx")
  if !ok {t.Error()}
}

func TestAbbConUnElementoObtener(t *testing.T) {
  a := abb.CrearAbb(cmpInt)
  a.Insertar(2, 2)
  ok := a.Obtener(2) == 2
  if !ok {t.Error()}
}

func TestAbbConVariosElementosObtener(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  ok := (a.Obtener("nil") == 21) &&
   (a.Obtener("hola") == "mundo") &&
    (a.Obtener("materia") == 8109)
  if !ok {t.Error()}

}

func TestAbbNoPerteneceObtenerEsNil(t *testing.T) {
  a := abb.CrearAbb(cmpInt)
  a.Insertar(1, 1)
  a.Insertar(3, 2)
  a.Insertar(2, 3)
  ok := a.Obtener(0) == nil
  if !ok {t.Error()}
}

func TestAbbConVariosElementosRemoverValorCorrecto(t *testing.T) {
  a := abb.CrearAbb(cmpInt)
  a.Insertar(7, "mundo")
  a.Insertar(4, 8109)
  a.Insertar(2, 21)
  a.Insertar(3, 15)
  a.Insertar(12, "h")
  a.Insertar(10, 15)
  a.Insertar(11, 21.12)
  a.Remover(10)
  removido := a.Remover(10)
  ok := removido == 15
  if !ok {t.Error()}
}

func TestAbbConVariosElementosRemoverCantidadReduceEnUno(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  a.Insertar("c", 15)
  a.Insertar("a", "h")
  a.Insertar("dd", 15)
  a.Insertar("da", 21.12)
  cant := a.Cantidad()
  a.Remover("a")
  ok := a.Cantidad() == (cant-1)
  if !ok {t.Error()}
}

func TestAbbConUnElementoRemoverValorCorrecto(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("hola", "mundo")
  removido := a.Remover("hola")
  ok := removido == "mundo"
  if !ok {t.Error()}
}

func TestAbbRemoverYNoPerteneceMas(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  a.Insertar("c", 15)
  a.Insertar("a", "h")
  a.Insertar("dd", 15)
  a.Insertar("da", 21.12)
  a.Remover("c")
  ok := !a.Pertenece("c")
  if !ok {t.Error()}
}

func TestRemoverYObtenerDevuelveNil(t *testing.T) {
  a := abb.CrearAbb(cmpString)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  a.Insertar("c", 15)
  a.Insertar("a", "h")
  a.Insertar("dd", 15)
  a.Insertar("da", 21.12)
  a.Remover("c")
  ok := a.Obtener("c") == nil
  if !ok {t.Error()}
}
