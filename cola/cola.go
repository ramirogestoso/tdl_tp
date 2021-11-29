package cola

type Cola struct {
  inicio *Nodo
  fin *Nodo
}

type Nodo struct {
  valor interface{}
  proximo *Nodo
}

func CrearCola() *Cola {
  return &Cola{nil, nil}
}

func (c *Cola) Encolar(x interface{}) {
  nodo := &Nodo{x, nil}
  if c.inicio == nil {
    c.inicio = nodo
  } else {
    c.fin.proximo = nodo
  }
  c.fin = nodo
}

func (c *Cola) Desencolar() interface{} {
  if c.inicio == nil {
    return nil
  }
  valorDesencolado := c.inicio.valor
  c.inicio = c.inicio.proximo
  return valorDesencolado
}

func (c *Cola) VerTope() interface{} {
  if c.inicio == nil {
    return nil
  }
  return c.inicio.valor
}

func (c *Cola) EstaVacia() bool {
  return c.inicio == nil
}
