package calculadora

import (
  "math"
)

type Operador interface {
  Calcular([]int) (int, bool)
  CantidadDeOperandos() int
}

type Suma struct{}
type Resta struct{}
type Multiplicacion struct{}
type Division struct{}
type Potencia struct{}
type Logaritmo struct{}
type Raiz struct{}
type Ternario struct{}

func (s Suma) Calcular(nums []int) (int, bool) {
  if len(nums) != 2 { return 0, false }
  suma := nums[0] + nums[1]
  return suma, true
}
func (s Suma) CantidadDeOperandos() int {
  return 2
}

func (r Resta) Calcular(nums []int) (int, bool) {
  if len(nums) != 2 { return 0, false }
  resta := nums[0] - nums[1]
  return resta, true
}
func (r Resta) CantidadDeOperandos() int {
  return 2
}

func (m Multiplicacion) Calcular(nums []int) (int, bool) {
  if len(nums) != 2 { return 0, false }
  mult := nums[0] * nums[1]
  return mult, true
}
func (m Multiplicacion) CantidadDeOperandos() int {
  return 2
}

func (d Division) Calcular(nums []int) (int, bool) {
  if len(nums) != 2 || nums[1] == 0 { return 0, false }
  div := int(nums[0] / nums[1])
  return div, true
}
func (d Division) CantidadDeOperandos() int {
  return 2
}

func (p Potencia) Calcular(nums []int) (int, bool) {
  if len(nums) != 2 || nums[1] < 0 { return 0, false }
  pot := int(math.Pow(float64(nums[0]), float64(nums[1])))
  return pot, true
}
func (p Potencia) CantidadDeOperandos() int {
  return 2
}

func (l Logaritmo) Calcular(nums []int) (int, bool) {
  if len(nums) != 2 || nums[1] < 2 { return 0, false }
  log := int(math.Log10(float64(nums[0])) / math.Log10(float64(nums[1])))
  return log, true
}
func (l Logaritmo) CantidadDeOperandos() int {
  return 2
}

func (r Raiz) Calcular(nums []int) (int, bool) {
  if len(nums) != 1 || nums[0] < 0 { return 0, false }
  raiz := int(math.Sqrt(float64(nums[0])))
  return raiz, true
}
func (r Raiz) CantidadDeOperandos() int {
  return 1
}

func (t Ternario) Calcular(nums []int) (int, bool) {
  if len(nums) != 3 { return 0, false }
  tern := nums[1]
  if nums[0] == 0 { tern = nums[2] }
  return tern, true
}
func (t Ternario) CantidadDeOperandos() int {
  return 3
}
