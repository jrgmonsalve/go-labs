package main

import (
	"fmt"
)

func main() {
	canal := make(chan int) // Crea un canal para int.
	go generador(canal)     // Inicia la goroutine que actúa como un generador.

	// Recibe valores del canal, similar a iterar sobre los valores generados por 'yield' en Python.
	for valor := range canal {
		fmt.Println(valor)
	}
}

func generador(canal chan int) {
	for i := 0; i < 1000000; i++ {
		canal <- i // Envía el valor de i al canal, similar a 'yield' en Python.
	}
	close(canal) // Cierra el canal después de enviar todos los valores.
}

func yieldExerices() {
	for i := 0; i < 1000000; i++ {
		fmt.Println(i)
	}
}

func makeExcercises() {
	emails := make([]string, 0, 4)
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(cap(emails))
}

type Email struct {
	Subject string
	Body    string
}

func mapsExercises() {
	urls := map[string]string{
		"google":   "http://google.com",
		"facebook": "http://facebook.com",
	}
	fmt.Println(urls)
	urls["twitter"] = "http://twitter.com"
	fmt.Println(urls)
	delete(urls, "facebook")
	fmt.Println(urls)
}
func mergeSlices() {
	users1 := []string{"user1", "user2", "user3"}
	fmt.Println(users1)
	users2 := []string{"user4", "user5", "user6"}
	fmt.Println(users2)
	users1 = append(users1, users2...)
	fmt.Println(users1)
}

func slicesExerices() {
	prices := []float64{1.99, 2.99, 3.99, 4.99}
	fmt.Println(prices)
	fmt.Println(prices[1:3])
	fmt.Println(prices[:3])
	fmt.Println(prices[0:3]) // slice with the first three elements
	fmt.Println(prices[1:])  // slice with all elements starting from the second
	fmt.Println(prices[:1])  // slice with one element starting from the beginning
	fmt.Println(prices[:0])  // empty slice
	fmt.Println(prices[:])
	prices = append(prices, 5.99)
	fmt.Println(prices)
	fmt.Println(prices[len(prices)-1:])
	fmt.Println(prices[len(prices)-1])
}

func arraysExerices() {
	// Arrays
	prices := [4]float64{}
	prices[0] = 1.99
	fmt.Println(prices)
	fmt.Println(len(prices))
	fmt.Println(cap(prices))

	var productNames [4]string = [4]string{"carrot", "apple", "banana", "orange"}
	fmt.Println(productNames)

	featurePrices := prices[1:3] // take a slice of the array
	fmt.Println(featurePrices)
	featurePrices[0] = 2.0
	allPricesExceptFirst := prices[1:] // take a slice of the array
	fmt.Println(allPricesExceptFirst)

	lastPrices := prices[len(prices)-1:]
	fmt.Println(lastPrices)
	fmt.Println(prices)
}
