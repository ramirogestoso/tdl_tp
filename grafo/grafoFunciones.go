package grafo

import (
  "github.com/ramirogestoso/tp/cola"
)

func EsConexo(grafo *Grafo) bool {
  visitados := make(map[interface{}]bool)
  v := grafo.VerticeAleatorio()
  visitarAdyacentes(grafo, visitados, v)
  for _, w := range grafo.Vertices() {
    _, visitado := visitados[w]
    if !visitado { return false }
  }
  return true
}

func visitarAdyacentes(grafo *Grafo, visitados map[interface{}]bool, v interface{}) {
  if _, visitado := visitados[v]; visitado == true { return }
  visitados[v] = true
  for _, w := range grafo.Adyacentes(v) {
    visitarAdyacentes(grafo, visitados, w)
  }
}

func nuevoColor(i int) int {
  if i == 1 { return 0 }
  return 1
}

func EsBipartito(grafo *Grafo) bool {
  padres := make(map[interface{}]interface{})
  for _, v := range grafo.Vertices() {
    _, visitado := padres[v]
    if !visitado {
      if !_esBipartito(grafo, padres, v) { return false }
    }
  }
  return true
}

func _esBipartito(grafo *Grafo, padres map[interface{}]interface{}, vi interface{}) bool {
  color := make(map[interface{}]int)
  cola := cola.CrearCola()
  padres[vi] = nil
  color[vi] = 0
  cola.Encolar(vi)
  for !cola.EstaVacia() {
    v := cola.Desencolar()
    for _, w := range grafo.Adyacentes(v) {
      _, fueVisitadoW := padres[w]
      if !fueVisitadoW {
        padres[w] = v
        color[w] = nuevoColor(color[v])
        cola.Encolar(w)
      } else { // ya fue visitado por otro padre
        if color[w] == color[v] { return false }
      }
    }
  }
  return true
}

func Aciclico(grafo *Grafo) bool {
  padres := make(map[interface{}]interface{})
  for _, v := range grafo.Vertices() {
    _, visitado := padres[v]
    if !visitado {
      if !_esAciclico(grafo, padres, v) { return false }
    }
  }
  return true
}

func _esAciclico(grafo *Grafo, padres map[interface{}]interface{}, vi interface{}) bool {
  cola := cola.CrearCola()
  cola.Encolar(vi)
  padres[vi] = nil
  for !cola.EstaVacia() {
    v := cola.Desencolar()
    for _, w := range grafo.Adyacentes(v) {
      _, fueVisitadoW := padres[w]
      if !fueVisitadoW {
        padres[w] = v
        cola.Encolar(w)
      } else {
        if w != padres[v] { return false }
      }
    }
  }
  return true
}
