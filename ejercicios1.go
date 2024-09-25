package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// var profesion = "Programador de software"


func main() {
	// salario := 1000000
	// fmt.Println("Profesion: ", profesion)
	// fmt.Println("Sueldo: ", salario)
	// fmt.Println("texto desde la función main")
	// imprimir()
	// print(hola_string("Fabián"))
	// fmt.Println(sumar(3, 5))
	// comparar_bool()
	// arreglo()
	// tipos_datos()
	// convertStringToBoolean()
	// convertBooleanToString()
	// ejercicio_15()
	// ejercicio_15_2()
	// ejercicio_16()
	// ejercicio_17()
	// ejercicio_18()
	// ejercicio_19()
	
	fmt.Println("The height is: ", height, "meters")
	fmt.Println("The height is: ", equivalenceInFeet(height), "feet")
	fmt.Println("The new height is: ", height, "meters") 
	// ejercicio_22()
	fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
}

func imprimir() {
	fmt.Println("Texto desde la función imprimir")
}

func hola_string(s string) string {
	return "Hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("A y B son iguales")
	} else {
		fmt.Println("A y B no son iguales")
	}
}

func arreglo() {
	var aprendices [5] string
	
	aprendices[0] = "Leonardo,"
	aprendices[1] = "Juan Sebastián,"
	aprendices[2] = "Frank,"
	aprendices[3] = "Juan José,"
	aprendices[4] = "Jaider,"

	fmt.Println(aprendices[0])
}

func tipos_datos() {
	var texto string = "Fabián"
	var numero int = 1000
	var decimal float64 = 0.0001

	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func convertStringToBoolean () {
	var palabra string = "true"

	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("Este es el error", err)
}

func convertBooleanToString () {
	var palabra_bool bool = true

	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func ejercicio_15() {
	var gravedad_tierra, gravedad_marte, gravedad_jupiter float32 = 9.8, 3.7, 24.8
	fmt.Println(gravedad_tierra, gravedad_marte, gravedad_jupiter)
}

func ejercicio_15_2() {
	var (
		Colombia bool = true
		Italia bool = false
		Islandia bool = true
	)
	fmt.Printf("Colombia  un país sudamericano? %t \n", Colombia)
	fmt.Printf("Italia es un país Asiático? %t \n", Italia)
	fmt.Printf("Islandia es un país europeo? %t \n", Islandia)
}

func ejercicio_16() {
	var Ciencia string
	var Contaduría int
	var Ingeniería float64
	var Derecho bool
	fmt.Println("Ciencia: ", Ciencia)
	fmt.Println("Contaduría: ", Contaduría)
	fmt.Println("Ingeniería: ", Ingeniería)
	fmt.Println("Derecho: ", Derecho)
}

func ejercicio_17() {
	nombre := "Benjamin Button"
	edad := 99
	peso := 80
	estudiante := false
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", estudiante)
}

var Gear1 = "Esta es el cambio 1"

func ejercicio_18() {
	var Gear2 = "Esta es el cambio 2"
	{
		var Gear3 = "Este es el cambio 3"
        fmt.Println(Gear3)
	}
	fmt.Println(Gear1)
	fmt.Println(Gear2)
}

func ejercicio_19 () {
	// color := "rojo"
	// fmt.Println(color, &color)

	vehicle1 := "red"
	fmt.Println("Vehicle1 is", vehicle1)

	vehicle2 := vehicle1
	fmt.Println("Vehicle2 is", vehicle2)

	vehicle3 := &vehicle1
	fmt.Println("Vehicle3 is", *vehicle3)

	vehicle1 = "grey"

	fmt.Println("Vehicle1 is", vehicle1, &vehicle1)
	fmt.Println("Vehicle2 is", vehicle2, &vehicle2)
	fmt.Println("Vehicle 3 is", *vehicle3, vehicle3)
}


func equivalenceInFeet(height float32) float32 {
	height = height * 3.28
	return height 
}

var height float32 = 1.70

const Pi = 3.1416

func area(radio float64) float64 {
	return Pi * radio * radio
}


// package main

// import ("fmt")

// func main() {
// 	// Hola mundo en Golang

// 	fmt.Println("Hola, Go!")

// 	/*
// 		Esto es un comentario en Go
// 	*/

// 	// https://go.dev/learn/
// }