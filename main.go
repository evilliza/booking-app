package main

import (
	"fmt"
	"math/rand"
)

func main() {

	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))
	for i := 0; i < 5; i++ {
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Printf("Math/Rand1: %d , Math/Rand2: %d\n", rnd1, rnd2)
			break
		}
	}

}
