package main 

import (
    "time"
    "github.com/stianeikeland/go-rpio"
)

func main() {
    _ = rpio.Open()

    pinNine := rpio.Pin(0)
    pinTen := rpio.Pin(10)

    pinNine.Input()
    pinTen.Input()
    count := 0
    for {
        println(pinNine.Read())
        time.Sleep(100 * time.Millisecond)
        count++

        if count >= 10 {
            rpio.Close()
            _ = rpio.Open()
        }
        /*
        if pinNine.Read() == rpio.High {
            break;
        } else {
            println(pinTen.Read())
            time.Sleep(100 * time.Millisecond)
        }
        */
    }

    rpio.Close()
}
