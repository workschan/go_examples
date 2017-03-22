// Get sqrt by Newton's method
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    v := 0.000001
    z := x
    for p := 0.0; math.Abs(p-z) >= v; {
        p = z
        z = z - (z*z - x)/(2*z)        
        
    }
    return z
}

func main() {
    fmt.Println(Sqrt(3))
    fmt.Println(math.Sqrt(3))
}

