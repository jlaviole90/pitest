package main 

import (
    "time"
    "github.com/stianeikeland/go-rpio"
)

func main() {
    err := rpio.Open()
    if err != nil {
        panic(err)
    }

    pinNine := rpio.Pin(0)
    pinTen := rpio.Pin(10)

    pinNine.Input()
    pinTen.Input()
    for {
        if pinNine.Read() == rpio.High {
            break;
        } else {
            println(pinTen.Read())
            time.Sleep(100 * time.Millisecond)
        }
    }

    rpio.Close()
}
