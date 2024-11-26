package main

import (
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
    if err := rpio.Open(); err != nil {
        println(err)
        os.Exit(1)
    }
    defer rpio.Close()

    pinNine := rpio.Pin(9)
    pinNine.Input()
    pinTen := rpio.Pin(10)
    pinTen.Input()

    for {
        println(pinNine.Read())
        println(pinTen.Read())
        time.Sleep(300 * time.Millisecond)

    }
}
