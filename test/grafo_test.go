package test

import (
  "testing"
  "../grafo"
)

func TestGrafoVacio(t *testing.T) {
  g := grafo.CrearGrafo(true)
  ok := g.Largo() == 0
  if !ok {t.Error()}
}

func TestAgregarVertices(t *testing.T) {
  g := grafo.CrearGrafo(true)
  g.AgregarVertice(1)
  g.AgregarVertice(2)
  g.AgregarVertice(3)
  ok := g.Largo() == 3
  if !ok {t.Error()}
}

func TestAgregarAristas(t *testing.T) {
  g := grafo.CrearGrafo(false)
  g.AgregarVertice(1)
  g.AgregarVertice(2)
  g.AgregarVertice(3)
  g.AgregarArista(1, 2)
  g.AgregarArista(1, 3)
  g.AgregarArista(2, 3)
  ok := g.EstanUnidos(1,2) && g.EstanUnidos(3, 1) && g.EstanUnidos(3, 2)
  if !ok {t.Error()}
}

func TestGrado(t *testing.T) {
  g := grafo.CrearGrafo(true)
  g.AgregarVertice("1")
  g.AgregarVertice(2)
  g.AgregarVertice(3)
  g.AgregarArista("1", 2)
  g.AgregarArista(3, "1")
  g.AgregarArista("1", 3)
  ok := g.Grado("1") == 2 && g.Grado(2) == 0 && g.Grado(3) == 1
  if !ok {t.Error()}
}

func TestGrafoConVerticesDeDistintoTipo(t *testing.T) {
  g := grafo.CrearGrafo(true)
  a := [2]int{1,2}
  b := make(chan int)
  c := 12
  d := "tdl"

  g.AgregarVertice(a)
  g.AgregarVertice(b)
  g.AgregarVertice(c)
  g.AgregarVertice(d)
  g.AgregarArista(a, b)
  g.AgregarArista(a, c)
  g.AgregarArista(a, d)
  g.AgregarArista(c, b)
  ok := !g.EstanUnidos(b,a) && g.EstanUnidos(a,c) && g.EstanUnidos(a,d) &&
    g.EstanUnidos(c,b) && !g.EstanUnidos(b,d) && !g.EstanUnidos(c,d)
  if !ok {t.Error()}
}

func TestVerticeQueNoExiste(t *testing.T) {
  g := grafo.CrearGrafo(true)
  g.AgregarVertice(0)
  ok := !g.EstanUnidos(1, 0)
  if !ok {t.Error()}
}


// hacer test para algoritmos de recorridos
// si es conexo, cantidad de componentes conexas (y fuertemente conexas para dirigidos)
// distancias, cantidad a distancia n, bipartito...
