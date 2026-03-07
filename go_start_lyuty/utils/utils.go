package utils

import (
	"fmt"
	"time"
)

var D = 4

func GetValue() string{
    return "Hello from this another package"
}

func Out(from, to int) {
  for i := from; i <= to; i++ {
    time.Sleep(50 * time.Millisecond)
    fmt.Println(i)
  }
}