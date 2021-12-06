package abb

//si clave1 > clave2 -> devuelve > 0
// si clave2 > clave1 -> devuelve < 0
// si clave1 == clave2 -> devuelve 0
type Cmp func(clave1 interface{}, clave2 interface{}) int

type Nodo struct {
  clave interface{}
  valor interface{}
  izq *Nodo
  der *Nodo
}

type Abb struct {
  raiz *Nodo
  cantidad int
  cmp Cmp
}

func CrearAbb(cmp Cmp) *Abb {
  return &Abb{nil, 0, cmp}
}

func (abb *Abb) Insertar(clave interface{}, valor interface{}) {
  abb.cantidad++
  nodo := &Nodo{clave, valor, nil, nil}
  if abb.raiz == nil {
    abb.raiz = nodo
    return
  }
  abb.raiz.insertar(nodo, abb.cmp)
}

func (actual *Nodo) insertar(nodo *Nodo, cmp Cmp) {

  switch comparacion := cmp(nodo.clave, actual.clave); {

  case comparacion < 0:
      if actual.izq != nil {
        actual.izq.insertar(nodo, cmp)
        } else { actual.izq = nodo }

  case comparacion > 0:
      if actual.der != nil {
        actual.der.insertar(nodo, cmp)
        } else { actual.der = nodo }

  default:
      actual.valor = nodo.valor

  }
}

func (abb *Abb) Remover(clave interface{}) interface{} {
  if abb.raiz == nil { return nil }
  nodo := abb.raiz.obtenerReferenciaANodo(clave, abb.cmp)
  if nodo == nil { return nil }
  nodoRemovido := *nodo
  valorRemovido := nodoRemovido.valor

  if nodoRemovido.izq != nil && nodoRemovido.der != nil { // hijos == 2
    claveDeReemplazante := nodoRemovido.obtenerClaveDeReemplazo()
    valorDeReemplazante := abb.Remover(claveDeReemplazante) // deberia tener 0 o 1 hijo
    (*nodo).clave = claveDeReemplazante
    (*nodo).valor = valorDeReemplazante
    return valorRemovido
  }
  // hijos < 2
  siguiente := nodoRemovido.izq
  if siguiente == nil { siguiente = nodoRemovido.der} // hijos == 0 si se cumple
  *nodo = siguiente
  abb.cantidad--
  return valorRemovido
}

func (nodo *Nodo) obtenerClaveDeReemplazo() interface{} {
  nodo = nodo.der
  for nodo.izq != nil { nodo = nodo.izq }
  return nodo.clave
}

func (actual *Nodo) obtenerReferenciaANodo(clave interface{}, cmp Cmp) **Nodo {
  switch comparacion := cmp(clave, actual.clave); {
  case comparacion < 0:
    if actual.izq == nil { return nil }
    return actual.izq.obtenerReferenciaANodo(clave, cmp)
  case comparacion > 0:
    if actual.der == nil { return nil }
    return actual.der.obtenerReferenciaANodo(clave, cmp)
  default:
    return &actual
  }
}

func (abb *Abb) Pertenece(clave interface{}) bool {
  if abb.raiz == nil { return false }
  return abb.raiz.pertenece(clave, abb.cmp)
}

func (actual *Nodo) pertenece(clave interface{}, cmp Cmp) bool {

  switch comparacion := cmp(clave, actual.clave); {

  case comparacion < 0:
    if actual.izq == nil { return false }
    return actual.izq.pertenece(clave, cmp)

  case comparacion > 0:
    if actual.der == nil { return false }
    return actual.der.pertenece(clave, cmp)

  default:
    return true
  }
}

func (abb *Abb) Obtener(clave interface{}) interface{} {
    if abb.raiz == nil { return nil }
    return abb.raiz.obtener(clave, abb.cmp)
}

func (actual *Nodo) obtener(clave interface{}, cmp Cmp) interface{} {

  switch comparacion := cmp(clave, actual.clave); {

  case comparacion < 0:
    if actual.izq == nil { return nil }
    return actual.izq.obtener(clave, cmp)

  case comparacion > 0:
    if actual.der == nil { return nil }
    return actual.der.obtener(clave, cmp)

  default:
    return actual.valor
  }
}

func (abb *Abb) Cantidad() int {
  return abb.cantidad
}
