// soal koding 1

// F(n) = [(n!) / (2^n)]

package main

import (
	"fmt"
	"math"
)

func hitungFactorial(n float64) float64 {
	factorial := 1
	for i := 1; i <= int(n); i++ {
		factorial *= i
	}

	return float64(factorial)

}

func hitungDuaPangkat(n float64) float64 {
	hasilDuaPangkat := math.Pow(2, n)

	return hasilDuaPangkat
}

func main() {
	n := 5.0

	hasilFn := hitungFactorial(n) / (hitungDuaPangkat(n))

	fmt.Println(hasilFn)

}
