package main

import (
	"log"
	"time"
)

func greeting(phrases string) {
	log.Println("Hello, ", phrases)
}

func slowGreeting(phrases string, c chan bool) {
	time.Sleep(3 * time.Second)
	log.Println("Hello, ", phrases)
	c <- true
}

func main() {
	//greeting("City")
	//greeting("World")
	c := make(chan bool)
	go slowGreeting("Universe", c)
	//greeting("me")
	<-c
}
