/* Execute : go run name.go */
package main

import (
	"fmt" 
	"strings"
)

/* global variable declaration */
var g int

func main(){
	fmt.Println("Hello, world!")
	fmt.Println("I'm in Go Programming World!")

	fmt.Println()

	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	fmt.Println()

	var y float64 = 20.0
	z := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)

	fmt.Println()

	var a, b, c = 3, 4, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

	fmt.Println()

	fmt.Printf("Hello\tWorld!\n")

	fmt.Println()

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)

	fmt.Println()

	/*
	for true {
		fmt.Printf("This loop will run forever.\n");
	}
	*/

	fmt.Println()

	/* calling a function to get max value */
	a = 100
	b = 200
	ret := max(a, b)
	fmt.Printf("Max value is : %d\n", ret)

	fmt.Println()

	d, e := swap("Mahesh", "Kumar")
	fmt.Println(d, e)

	fmt.Println()

	var greeting = "Hello world!"
	fmt.Printf("normal string : ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes : ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	fmt.Println()
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
	fmt.Printf("quoted string : ")
	fmt.Printf("%+q\n", sampleText)
	fmt.Printf("String Length is : %d\n", len(greeting))

	fmt.Println()
	
	//greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))

	fmt.Println()

	var n [10] int
	var i, j int
	/* initialize element of array n to 0 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	fmt.Println()

	var f int = 10
	fmt.Printf("Address of a variable : %x\n", &f)
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* returning multiple values from function */
func swap(x, y string) (string, string){
	return y, x
}
