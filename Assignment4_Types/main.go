package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := true, false
	fmt.Println("a =", a, "b =", b)
	c := a && b
	d := a || b
	fmt.Println("c =", c, "d =", d)

	var i int = -12000
	j := 12000
	fmt.Println("i:", i, "j:", j)
	fmt.Printf("The type of i is %T, the size of i is %d", i, unsafe.Sizeof(i))
	fmt.Printf("\nThe type of j is %T, the size of j is %d", j, unsafe.Sizeof(j))

	num1, num2 := 5.2, 8.3
	fmt.Printf("\ntype of num1 %T, num2 %T \n", num1, num2)
	sum := num1 + num2
	diff := num2 - num1
	fmt.Println("Sum:", sum, "Diff:", diff)
	num3, num4 := 59, 69
	fmt.Println("Sum:", num3+num4, "Diff:", num3-num4)

	c1 := complex(8, 7)
	c2 := 27 + 35i
	cadd := c1 + c2
	fmt.Println("c1:", c1, "c2:", c2, "sum:", cadd)
	cmul := c1 * c2 //27*8 + 27*7i + 8*35i -7*35 = -29 + 469i
	fmt.Println("Product: ", cmul)

	int1 := 10
	float1 := 67.8
	sum1 := int1 + int(float1)
	fmt.Println("Type Conversion: ", sum1)
	var float2 float64 = float64(int1)
	fmt.Println("float64 of int1:", float2)
}
