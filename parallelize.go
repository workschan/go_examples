package main

import "runtime"

type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += v[i]
    }
    c <- 1    // signal that this piece is done
}

var numCPU = runtime.NumCPU() // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}

func main() {
	v := Vector{}
	n := 1.0 * 10e7
	for i:=1.0; i<n; i++ {
		v = append(v, i)
	}
	
	v.DoAll(v)
	
}
