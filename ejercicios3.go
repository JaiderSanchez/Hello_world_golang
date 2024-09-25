package main

import (
    "fmt"
)

const Pi = 3.14159265

// Point 23 of the guide
func circle(radio float64) (area float64, perimeter float64, radio1 float64) {
	area = Pi * radio * radio
    perimeter = 2 * Pi * radio
	radio1 = radio
    return area, perimeter, radio1
}


// Point 24 of the guide

func add(numbers ... int) int {
	// The initial total is 0
	total := 0
	// Go through all the numbers
	for _, number := range numbers {
        // In each iteration add the value of the number to the total
		total = number + 1
    }
    // Return the total value
    return total
}

// Point 25 of the guide

func echoMountain(message string, iterations uint) {
	if iterations > 1 {
		echoMountain(message, iterations - 1)
	} 
	fmt.Println(message, iterations)	
	}


// Exercise 30

func ejercicio_else_if() {
	var toy string
	fmt.Println("Choose person, animal or thing: ")
	fmt.Scanln(&toy)
	if toy == "person" {
		fmt.Println("The object is a person")
	} else if toy == "thing" {
		fmt.Println("The object is a thing")
	} else if toy == "animal" {
		fmt.Println("The object is an animal")
    } else {
		fmt.Println("The object is another category")
	}
}

func ejercicio_residuo() {
	var x int
	var y int
	fmt.Println("Ingresa el valor de x: ")
	fmt.Scanln(&x)
	fmt.Println("Ingresa el valor de y: ")
	fmt.Scanln(&y)
	resultado := x / y
	fmt.Printf("El resultado de la división de %d y %d es: %d ", x, y, resultado)
	fmt.Printf("El residuo de la división de %d y %d es: %d ", x, y, x%y)
}

func ejercicio_par_impar() {
	var x int
	fmt.Println("Ingresa el valor de x: ")
	fmt.Scanln(&x)
	if x%2 == 0 {
        fmt.Println("El número", x, "es par")
    } else {
        fmt.Println("El número", x, "es impar")
    }
}

func main() {
	// a, p, r := circle(8)
	// fmt.Println("The area of ​​the circle is: ", a)
	// fmt.Println("The perimeter of the circle is: ", p)
	// fmt.Println("The radius of the circle is: ", r)

	// fmt.Println(add(2))
	// fmt.Println(add)
	// echoMountain("yodelayheehoo", 5)
	// ejercicio_else_if()
	// ejercicio_residuo()
	ejercicio_par_impar()
}