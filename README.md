WiringPi-Go
============

Golang wrapped version of Gordon's Arduino-like WiringPi for the Raspberry Pi

# Installation

install WiringPi first

```
cd WiringPi\wiringPi
sudo make install

go get github.com/hugozhu/rpi
```

# GPIO numbering

 ----------------------------------Model B2------------------------------------
 | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
 ------------------------------------------------------------------------------
 |     |     |    3.3v |      |   |  1 || 2  |   |      | 5v      |     |     |
 |   2 |   8 |   SDA.1 |   IN | 1 |  3 || 4  |   |      | 5V      |     |     |
 |   3 |   9 |   SCL.1 |   IN | 1 |  5 || 6  |   |      | 0v      |     |     |
 |   4 |   7 | GPIO. 7 |   IN | 1 |  7 || 8  | 1 | ALT0 | TxD     | 15  | 14  |
 |     |     |      0v |      |   |  9 || 10 | 1 | ALT0 | RxD     | 16  | 15  |
 |  17 |   0 | GPIO. 0 |  OUT | 1 | 11 || 12 | 0 | IN   | GPIO. 1 | 1   | 18  |
 |  27 |   2 | GPIO. 2 |   IN | 0 | 13 || 14 |   |      | 0v      |     |     |
 |  22 |   3 | GPIO. 3 |   IN | 0 | 15 || 16 | 0 | IN   | GPIO. 4 | 4   | 23  |
 |     |     |    3.3v |      |   | 17 || 18 | 0 | IN   | GPIO. 5 | 5   | 24  |
 |  10 |  12 |    MOSI |   IN | 0 | 19 || 20 |   |      | 0v      |     |     |
 |   9 |  13 |    MISO |   IN | 0 | 21 || 22 | 0 | IN   | GPIO. 6 | 6   | 25  |
 |  11 |  14 |    SCLK |   IN | 0 | 23 || 24 | 1 | IN   | CE0     | 10  | 8   |
 |     |     |      0v |      |   | 25 || 26 | 1 | IN   | CE1     | 11  | 7   |
 ------------------------------------------------------------------------------
 |  28 |  17 | GPIO.17 |   IN | 0 | 51 || 52 | 0 | IN   | GPIO.18 | 18  | 29  |
 |  30 |  19 | GPIO.19 |   IN | 0 | 53 || 54 | 0 | IN   | GPIO.20 | 20  | 31  |
 ------------------------------------------------------------------------------
 | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
 -----------------------------------Model B2-----------------------------------

more to read at: [http://hugozhu.myalert.info/2013/03/22/19-raspberry-pi-gpio-port-naming.html](http://hugozhu.myalert.info/2013/03/22/19-raspberry-pi-gpio-port-naming.html)

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
go install github.com/hugozhu/rpi 
go run src/lcd.go 
```
