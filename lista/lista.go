package lista

type Nodo struct {
  valor interface{}
  anterior *Nodo
  proximo *Nodo
}

type Lista struct {
  largo int
  inicio *Nodo
  fin *Nodo
}

type Iterador struct {
  puedeAvanzar bool
  puedeRetroceder bool
  lista *Lista
  actual *Nodo
}

func CrearLista() *Lista {
  return &Lista{0, nil, nil}
}

func (l *Lista) EstaVacia() bool {
  return l.fin == nil
}

func (l *Lista) InsertarUltimo(x interface{}) {
  nodo := &Nodo{x, l.fin, nil}
  if l.inicio == nil { l.inicio = nodo } else { l.fin.proximo = nodo }
  l.fin = nodo
  l.largo++
}

func (l *Lista) InsertarPrimero(x interface{}) {
  if l.EstaVacia() {
    l.InsertarUltimo(x)
  } else {
    nodo := &Nodo{x, nil, l.inicio}
    l.inicio.anterior = nodo
    l.inicio = nodo
    l.largo++
  }
}

func (l *Lista) InsertarEnPosicion(i int, x interface{}) {
  if i == l.largo {
    l.InsertarUltimo(x)
  } else if i == 0 {
    l.InsertarPrimero(x)
  } else {
    nodo := l._ObtenerNodoEnPosicion(i)
    nodoNuevo := &Nodo{x, nodo.anterior, nodo}
    nodo.anterior.proximo = nodoNuevo
    nodo.anterior = nodoNuevo
    l.largo++
  }
}

func (l *Lista) ObtenerEnPosicion(i int) interface{} {
  nodo := l._ObtenerNodoEnPosicion(i)
  if nodo == nil { return nil }
  return nodo.valor
}

func (l *Lista) ObtenerUltimo() interface{} {
  if l.fin == nil { return nil }
  return l.fin.valor
}

func (l *Lista) ObtenerLargo() int {
  return l.largo
}

func (l *Lista) _ObtenerNodoEnPosicion(i int) *Nodo {
  if i < 0 || i >= l.largo { return nil }
  nodo := l.inicio
  if i <= l.largo/2 {
    for j:=0; j<i; j++ {
      nodo = nodo.proximo
    }
  } else {
    nodo = l.fin
    for j:=l.largo-1; j>i; j-- {
      nodo = nodo.anterior
    }
  }
  return nodo
}

func (l *Lista) RemoverEnPosicion(i int) interface{} {
  if i == l.largo-1 {
    return l.RemoverUltimo()
  }
  nodo := l._ObtenerNodoEnPosicion(i)
  nodo.anterior.proximo = nodo.proximo
  nodo.proximo.anterior = nodo.anterior
  l.largo--
  return nodo.valor
}

func (l *Lista) RemoverUltimo() interface{} {
  if l.fin == nil { return nil }
  nodo := l.fin
  l.fin = nodo.anterior
  l.fin.proximo = nil
  l.largo--
  return nodo.valor
}

func (l *Lista) RemoverPrimero() interface{} {
  if l.inicio == nil { return nil }
  nodo := l.inicio
  l.inicio = nodo.proximo
  l.inicio.anterior = nil
  l.largo--
  return nodo.valor
}

func (l *Lista) Encontrar(x interface{}) int {
  i := 0
  nodo := l.inicio
  for nodo != nil {
    if nodo.valor == x {
      return i
    }
    nodo = nodo.proximo
    i++
  }
  return -1
}

func (l *Lista) CrearIterador() *Iterador {
  desplazarse := !l.EstaVacia()
  return &Iterador{desplazarse, desplazarse, l, l.inicio}
}

func (it *Iterador) Avanzar() bool {
  if !it.puedeAvanzar { return false }
  if it.actual != nil {
    it.actual = it.actual.proximo
    it.puedeAvanzar = it.actual != nil
  } else { // actual == nil y esta al inicio
    it.actual = it.lista.inicio
    it.puedeRetroceder = true
  }
  return it.puedeAvanzar
}

func (it *Iterador) Retroceder() bool {
  if !it.puedeRetroceder { return false }
  if it.actual != nil {
    it.actual = it.actual.anterior
    it.puedeRetroceder = it.actual != nil
  } else { // actual == nil y esta al final
    it.actual = it.lista.fin
    it.puedeAvanzar = true
  }
  return it.puedeRetroceder
}

func (it *Iterador) VerActual() interface{} {
  if it.actual == nil { return nil }
  return it.actual.valor
}
