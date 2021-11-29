package test

import (
  "testing"
  "../abb"
)

func cmp(c1 string, c2 string) int {
  switch {
  case c2>c1: return -1
  case c1>c2: return 1
  default: return 0
  }
}

func TestAbbNuevoTieneCantidadCero(t *testing.T) {
  a := abb.CrearAbb(cmp)
  ok := a.Cantidad() == 0
  if !ok {t.Error()}
}

func TestAbbInsertarUnElementoCantidadEsUno(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("uno", 1)
  ok := a.Cantidad() == 1
  if !ok {t.Error()}
}

func TestAbbInsertarVariosCantidadCorrecta(t *testing.T) {
  a := abb.CrearAbb(cmp)
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
  a := abb.CrearAbb(cmp)
  a.Insertar("Hola Mundo", 12)
  ok := a.Pertenece("Hola Mundo")
  if !ok {t.Error()}
}

func TestAbbConVariosElementosPertenecen(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("uno", 1)
  a.Insertar("dos", 2)
  a.Insertar("tres", 3)
  ok := a.Pertenece("uno") && a.Pertenece("dos") && a.Pertenece("tres")
  if !ok {t.Error()}
}

func TestAbbConVariosElementosNoPertenece(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("uno", 1)
  a.Insertar("dos", 2)
  a.Insertar("tres", 3)
  ok := !a.Pertenece("unx")
  if !ok {t.Error()}
}

func TestAbbConUnElementoObtener(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("uno", 2)
  ok := a.Obtener("uno") == 2
  if !ok {t.Error()}
}

func TestAbbConVariosElementosObtener(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  ok := (a.Obtener("nil") == 21) &&
   (a.Obtener("hola") == "mundo") &&
    (a.Obtener("materia") == 8109)
  if !ok {t.Error()}

}

func TestAbbNoPerteneceObtenerEsNil(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("uno", 1)
  a.Insertar("dos", 2)
  a.Insertar("tres", 3)
  ok := a.Obtener("unx") == nil
  if !ok {t.Error()}
}

func TestAbbConVariosElementosRemoverValorCorrecto(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  a.Insertar("c", 15)
  a.Insertar("a", "h")
  a.Insertar("dd", 15)
  a.Insertar("da", 21.12)
  removido := a.Remover("c")
  ok := removido == 15
  if !ok {t.Error()}
}

func TestAbbConVariosElementosRemoverCantidadReduceEnUno(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("hola", "mundo")
  a.Insertar("materia", 8109)
  a.Insertar("nil", 21)
  a.Insertar("c", 15)
  a.Insertar("a", "h")
  a.Insertar("dd", 15)
  a.Insertar("da", 21.12)
  cant := a.Cantidad()
  a.Remover("c")
  ok := a.Cantidad() == (cant-1)
  if !ok {t.Error()}
}

func TestAbbConUnElementoRemoverValorCorrecto(t *testing.T) {
  a := abb.CrearAbb(cmp)
  a.Insertar("hola", "mundo")
  removido := a.Remover("hola")
  ok := removido == "mundo"
  if !ok {t.Error()}
}

func TestAbbRemoverYNoPerteneceMas(t *testing.T) {
  a := abb.CrearAbb(cmp)
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
  a := abb.CrearAbb(cmp)
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
