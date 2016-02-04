package main

import (
	"fmt"
  "math"
)

func main() {
  fmt.Println(Uint8FromInt(255))
}

func Uint8FromInt(x int) (uint8, error) {
  if 0 <= x && x <= math.MaxUint8 {
    return uint8(x), nil
  }
  return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
