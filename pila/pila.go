package pila

type Nodo struct {
  valor interface{}
  proximo *Nodo
}

type Pila struct {
  tope *Nodo
}

func CrearPila() *Pila {
  return &Pila{nil}
}

func (p *Pila) Apilar(x interface{}) {
  nodo := &Nodo{x, p.tope}
  p.tope = nodo
}

func (p *Pila) VerTope() interface{} {
  if p.tope == nil {
    return nil
  }
  return p.tope.valor
}

func (p *Pila) Desapilar() interface{} {
  if p.tope == nil {
    return nil
  }
  desapilado := p.tope.valor
  p.tope = p.tope.proximo
  return desapilado
}

func (p *Pila) EstaVacia() bool {
  return p.tope == nil
}
