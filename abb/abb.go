package abb

type Cmp func(clave1 string, clave2 string) int

type Nodo struct {
  clave string
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

func (abb *Abb) Insertar(clave string, valor interface{}) {
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

  case comparacion > 0:
      if actual.izq != nil {
        actual.izq.insertar(nodo, cmp)
        } else { actual.izq = nodo }

  case comparacion < 0:
      if actual.der != nil {
        actual.der.insertar(nodo, cmp)
        } else { actual.der = nodo }

  default:
      actual.valor = nodo.valor

  }
}

func (abb *Abb) Remover(clave string) interface{} {
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

func (nodo *Nodo) obtenerClaveDeReemplazo() string {
  nodo = nodo.der
  for nodo.izq != nil { nodo = nodo.izq }
  return nodo.clave
}

func (actual *Nodo) obtenerReferenciaANodo(clave string, cmp Cmp) **Nodo {
  switch comparacion := cmp(clave, actual.clave); {
  case comparacion > 0:
    if actual.izq == nil { return nil }
    return actual.izq.obtenerReferenciaANodo(clave, cmp)
  case comparacion < 0:
    if actual.der == nil { return nil }
    return actual.der.obtenerReferenciaANodo(clave, cmp)
  default:
    return &actual
  }
}

func (abb *Abb) Pertenece(clave string) bool {
  if abb.raiz == nil { return false }
  return abb.raiz.pertenece(clave, abb.cmp)
}

func (actual *Nodo) pertenece(clave string, cmp Cmp) bool {

  switch comparacion := cmp(clave, actual.clave); {

  case comparacion > 0:
    if actual.izq == nil { return false }
    return actual.izq.pertenece(clave, cmp)

  case comparacion < 0:
    if actual.der == nil { return false }
    return actual.der.pertenece(clave, cmp)

  default:
    return true
  }
}

func (abb *Abb) Obtener(clave string) interface{} {
    if abb.raiz == nil { return nil }
    return abb.raiz.obtener(clave, abb.cmp)
}

func (actual *Nodo) obtener(clave string, cmp Cmp) interface{} {

  switch comparacion := cmp(clave, actual.clave); {

  case comparacion > 0:
    if actual.izq == nil { return nil }
    return actual.izq.obtener(clave, cmp)

  case comparacion < 0:
    if actual.der == nil { return nil }
    return actual.der.obtener(clave, cmp)

  default:
    return actual.valor
  }
}

func (abb *Abb) Cantidad() int {
  return abb.cantidad
}
