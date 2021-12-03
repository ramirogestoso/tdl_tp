package grafo

type Grafo struct {
  dirigido bool
  grafo map[interface{}]map[interface{}]bool
}

func CrearGrafo(dirigido bool) *Grafo {
  g := make(map[interface{}]map[interface{}]bool)
  return &Grafo{grafo: g, dirigido: dirigido}
}

func (g *Grafo) AgregarVertice(x interface{}) {
  g.grafo[x] = make(map[interface{}]bool)
}

func (g *Grafo) AgregarArista(v1 interface{}, v2 interface{}) {
  // panic si no existe v1 o v2
  g.grafo[v1][v2] = true
  if !g.dirigido {
    g.grafo[v2][v1] = true
  }
}

func (g *Grafo) ExisteVertice(v interface{}) bool {
  _, ok := g.grafo[v]
  return ok
}

func (g *Grafo) EstanUnidos(v1 interface{}, v2 interface{}) bool {
  _, ok := g.grafo[v1][v2] // si v1 o v2 no existe, devuelve false
  return ok
}

func (g *Grafo) Grado(v interface{}) int {
  return len(g.grafo[v])
}

func (g *Grafo) Largo() int {
  return len(g.grafo)
}

// obtener vertices, adyacentes...
