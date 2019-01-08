package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := 1.0
  z := float64(1)

  while x - z > 0.1:
    z -= (z * z - x) / (2 * z)
}

func main() {
  fmt.Println(Sqrt(2))
}
