WiringPi-Go
============

Golang wrapped version of Gordon's Arduino-like WiringPi for the Raspberry Pi

# Installation

install WiringPi first

```
git clone git://git.drogon.net/wiringPi


cd WiringPi\wiringPi
sudo make install

go get github.com/hugozhu/rpi
```

# GPIO numbering

![gpio](http://wiringpi.com/wp-content/uploads/2013/03/gpio1.png)

more to read at: [http://wiringpi.com/pins/](http://wiringpi.com/pins/)

# Sample codes

## `lcd.go`
```go
package main

import (
    . "github.com/CarlosRA97/wpi"
)

func main() {
    WiringPiSetup()

    //use default pin naming
    PinMode(PIN_GPIO_4, OUTPUT)
    DigitalWrite(PIN_GPIO_4, LOW)
    Delay(400)
    DigitalWrite(PIN_GPIO_4, HIGH)

    //use raspberry pi board pin numbering, similiar to RPi.GPIO.setmode(RPi.GPIO.BOARD)
    Delay(400)
    DigitalWrite(BoardToPin(16), LOW)
    Delay(400)
    DigitalWrite(BoardToPin(16), HIGH)

    //use raspberry pi bcm gpio numbering, similiar to RPi.GPIO.setmode(RPi.GPIO.BCM)
    Delay(400)
    DigitalWrite(GpioToPin(23), LOW)
    Delay(400)
    DigitalWrite(GpioToPin(23), HIGH)
}
```

## Run

```
export GOPATH=`pwd`
go install github.com/CarlosRA97/wpi 
go run src/lcd.go 
```
