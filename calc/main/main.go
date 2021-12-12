package main

import (
  "github.com/ramirogestoso/tp/calc/calculadora"
  "os"
  "bufio"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
     calculadora.CalcularLinea(scanner.Text())
   }
}
