package calculadora

import (
  "fmt"
  "strings"
  "github.com/ramirogestoso/tp/pila"
  "strconv"
)

const (
  SUMA = "+"
  RESTA = "-"
  MULT = "*"
  DIV = "/"
  POTENCIA = "^"
  RAIZ = "sqrt"
  LOG = "log"
  TERN = "?"
  ERROR = "ERROR"
)

func CalcularLinea(linea string) {
  pilaNumeros := pila.CrearPila()
  elementos := strings.Split(linea, " ")

  for _, elemento := range elementos {
    num, esNum := strToInt(elemento)
    operador, esOp := strToOp(elemento)

    if esNum {
      pilaNumeros.Apilar(num)
    } else if esOp {
      resultado, ok := operar(operador, pilaNumeros)
      if !ok { // falla al operar
        fmt.Println(ERROR)
        return
      }
      pilaNumeros.Apilar(resultado)
    } else { // no es ninguna y la linea esta mal
      fmt.Println(ERROR)
      return
    }
  }
  resultadoFinal := pilaNumeros.Desapilar()
  if !pilaNumeros.EstaVacia() {
    fmt.Println(ERROR)
    return
  }
  fmt.Println(resultadoFinal)
}

func strToInt(s string) (int, bool) {
  num, err := strconv.Atoi(s)
  if err != nil { return 0, false }
  return num, true
}

func strToOp(elem string) (Operador, bool) {
  switch elem {
  case SUMA:
    return Suma{}, true
  case RESTA:
    return Resta{}, true
  case MULT:
    return Multiplicacion{}, true
  case DIV:
    return Division{}, true
  case POTENCIA:
    return Potencia{}, true
  case RAIZ:
    return Raiz{}, true
  case LOG:
    return Logaritmo{}, true
  case TERN:
    return Ternario{}, true
  }
  return nil, false
}

func operar(operador Operador, pilaNumeros *pila.Pila) (int, bool) {
  cantOperandos := operador.CantidadDeOperandos()
  nums := make([]int, cantOperandos)
  for i:=0; i<cantOperandos; i++ {
    if pilaNumeros.EstaVacia() { return 0, false }
    nums[cantOperandos-i-1] = pilaNumeros.Desapilar().(int)
  }
  return operador.Calcular(nums)
}
