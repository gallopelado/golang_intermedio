package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	// dos versiones del mismo paquete
	hellomod.SayHello()
	hellomod2.SayHello("holis")
}
