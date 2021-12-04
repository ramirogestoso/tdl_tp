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

func (g *Grafo) AgregarArista(v interface{}, w interface{}) {
  // panic si no existe v o w
  g.grafo[v][w] = true
  if !g.dirigido { g.grafo[w][v] = true }
}

func (g *Grafo) EliminarArista(v interface{}, w interface{}) {
  delete(g.grafo[v], w)
  if !g.dirigido { delete(g.grafo[w], v)}
}

func (g *Grafo) ExisteVertice(v interface{}) bool {
  _, ok := g.grafo[v]
  return ok
}

func (g *Grafo) EstanUnidos(v interface{}, w interface{}) bool {
  _, ok := g.grafo[v][w] // si v o w no existe, devuelve false
  return ok
}

func (g *Grafo) Grado(v interface{}) int {
  return len(g.grafo[v])
}

func (g *Grafo) Largo() int {
  return len(g.grafo)
}

func (g *Grafo) Vertices() []interface{} {
  // devuelve un slice con los vertices
  vertices := make([]interface{}, len(g.grafo))
  i:=0
  for k := range g.grafo {
    vertices[i] = k
    i++
  }
  return vertices
}

func (g *Grafo) Adyacentes(v interface{}) []interface{} {
  _, existe := g.grafo[v]
  if !existe { return nil }
  adyacentes := make([]interface{}, len(g.grafo[v]))
  i:=0
  for k := range g.grafo[v] {
    adyacentes[i] = k
    i++
  }
  return adyacentes
}
