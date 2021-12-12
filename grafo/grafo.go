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

func (g *Grafo) EliminarVertice(v interface{}) bool {
  if !g.ExisteVertice(v) { return false }
  delete(g.grafo, v)
  return true
}

func (g *Grafo) AgregarArista(v interface{}, w interface{}) bool {
  if !g.ExistenVertices(v, w){ return false }
  g.grafo[v][w] = true
  if !g.dirigido { g.grafo[w][v] = true }
  return true
}

func (g *Grafo) EliminarArista(v interface{}, w interface{}) bool {
  if !g.ExistenVertices(v, w){ return false }
  delete(g.grafo[v], w)
  if !g.dirigido { delete(g.grafo[w], v)}
  return true
}

func (g *Grafo) ExisteVertice(v interface{}) bool {
  _, ok := g.grafo[v]
  return ok
}

func (g *Grafo) ExistenVertices(vertices ...interface{}) bool {
  for _, v := range vertices {
    if _, ok := g.grafo[v]; !ok { return false }
  }
  return true
}

func (g *Grafo) EstanUnidos(v interface{}, w interface{}) bool {
  if !g.ExistenVertices(v, w){ return false }
  _, ok := g.grafo[v][w]
  return ok
}

func (g *Grafo) Grado(v interface{}) int {
  if !g.ExisteVertice(v) { return -1 }
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
  // devuelve in slice con los vertices adyacentes de v si existe, nil si no
  if !g.ExisteVertice(v) { return nil }
  adyacentes := make([]interface{}, len(g.grafo[v]))
  i:=0
  for k := range g.grafo[v] {
    adyacentes[i] = k
    i++
  }
  return adyacentes
}

func (g *Grafo) VerticeAleatorio() interface{} {
  var v interface{}
  for w,_ := range g.grafo {
    v = w
    break
  }
  return v
}
