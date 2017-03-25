package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if(x < 0) {
        return 0, ErrNegativeSqrt(x)
	}
    v := 0.000001
    z := x
    for p := 0.0; math.Abs(p-z) >= v; {
        p = z
        z = z - (z*z - x)/(2*z)      
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}

