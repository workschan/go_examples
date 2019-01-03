package main

import (
  "fmt"
  "strings"
)

func main() {
  const n = 99999
  const steps = 50
  const size = n/steps
  for i := 0; i < n+1; i++ {
    h := strings.Repeat("=", i/size) + strings.Repeat(" ", steps-i/size)
    fmt.Printf("\r%.0f%%[%s]", float64(i)/n*100, h)
  
  }
}
