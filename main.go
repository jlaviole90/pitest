package main 

import (
    "github.com/stianeikeland/go-rpio"
)

func main() {
    pin := rpio.Pin(10)

    pin.Output()
    println(pin.Read())

    pin.High()
    println(pin.Read())

    pin.Low()
    println(pin.Read())

    pin.Toggle()
    println(pin.Read())
}
