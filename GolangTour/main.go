package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add (x int, y int) int{
	return x + y
}

func swap (x, y string) (string, string) {
	return y, x
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main () {
	fmt.Println("The time now is", time.Now())
	fmt.Println("My favourite number is", rand.Int())
	fmt.Printf("Now you have %g problem.\n", math.Sqrt(7))
	fmt.Println(add(42, 13))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(split(17))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
