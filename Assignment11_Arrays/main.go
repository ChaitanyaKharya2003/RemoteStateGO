package main

import (
	"fmt"
	"math/rand"
)

func change0(arr [5]int) {
	arr[0] = 1
	fmt.Println("changing a4[0] to 1 inside function", arr)
}

func print32(arr [3][2]int) {
	for _, row := range arr {
		for _, item := range row {
			fmt.Printf(" %d", item)
		}
		fmt.Printf("\n")
	}

}
func print44(arr [4][4]float64) {
	for _, row := range arr {
		for _, item := range row {
			fmt.Printf(" %.2f", item)
		}
		fmt.Printf("\n")
	}

}

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := [3]string{"a", "b", "c"}
	a3 := a2
	a3[0] = "d"
	fmt.Println("a2 is", a2)
	fmt.Println("a3 is", a3)

	fmt.Println("a1 before change:", a1)
	for i := 0; i < len(a1); i++ {
		a1[i] = a1[i] + 2
	}
	fmt.Println("a1 after change:", a1)

	// passing to the function
	a4 := [...]int{6, 2, 3, 4, 5}
	fmt.Println("a4 before passing to function", a4)
	change0(a4)
	fmt.Println("a4 after passing to function", a4)

	a5 := [...]float64{1.2, 2.3, 3.4, 4.5}
	sum := float64(0)
	for i, v := range a5 {
		fmt.Printf("Index: %d, Value: %.2f\n", i+1, v)
		sum += v
	}
	fmt.Println("Sum of array a5:", sum)

	// multidimensional arrays
	m1 := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	print32(m1)

	var m2 [4][4]float64
	fmt.Println("Float array:")
	print44(m2)
	for i := 0; i < len(m2); i++ {
		for j := 0; j < len(m2[0]); j++ {
			m2[i][j] = rand.Float64() * 10
		}
	}
	fmt.Println("Float array after change:")
	print44(m2)

	a6 := [5]int{4, 5, 6, 7, 8}
	var splicea6 []int = a6[1:4]
	fmt.Println("a6:", a6, "Splice of a6(idx 1 to 3):", splicea6)
}
