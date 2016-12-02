package wpi

/*

#cgo LDFLAGS: -lwiringPi
#include <wiringPi.h>
#include <stdio.h>

*/
import "C"

import (
	"errors"
)

// wiringPi modes
const (
	WPI_MODE_PINS          = iota // C.WPI_MODE_PINS
	WPI_MODE_GPIO          = iota // C.WPI_MODE_GPIO
	WPI_MODE_GPIO_SYS      = iota // C.WPI_MODE_GPIO_SYS
	WPI_MODE_PHYS          = iota // C.WPI_MODE_PHYS
	WPI_MODE_PIFACE        = iota // C.WPI_MODE_PIFACE
	WPI_MODE_UNINITIALISED = -1   // C.WPI_MODE_UNINITIALISED
)

// Pins (pin naming = wiringPi numeration)
const (
	PIN_GPIO_0  = iota // 0
	PIN_GPIO_1  = iota
	PIN_GPIO_2  = iota
	PIN_GPIO_3  = iota
	PIN_GPIO_4  = iota
	PIN_GPIO_5  = iota
	PIN_GPIO_6  = iota
	PIN_GPIO_7  = iota
	PIN_SDA     = iota
	PIN_SCL     = iota
	PIN_CE0     = iota
	PIN_CE1     = iota
	PIN_MOSI    = iota
	PIN_MOSO    = iota
	PIN_SCLK    = iota
	PIN_TXD     = iota
	PIN_RXD     = iota
	PIN_GPIO_8  = iota
	PIN_GPIO_9  = iota
	PIN_GPIO_10 = iota
	PIN_GPIO_11 = iota
)

// Pin modes
const (
	INPUT            = iota // C.INPUT
	OUTPUT           = iota // C.OUTPUT
	PWM_OUTPUT       = iota // C.PWM_OUTPUT
	GPIO_CLOCK       = iota // C.GPIO_CLOCK
	SOFT_PWM_OUTPUT  = iota // C.SOFT_PWM_OUTPUT
	SOFT_TONE_OUTPUT = iota // C.SOFT_TONE_OUTPUT
	PWM_TONE_OUTPUT  = iota // C.PWM_TONE_OUTPUT
)
const (
	LOW  = iota // C.LOW
	HIGH = iota // C.HIGH
)

// Pull up/down/none
const (
	PUD_OFF  = iota // C.PUD_OFF
	PUD_DOWN = iota // C.PUD_DOWN
	PUD_UP   = iota // C.PUD_UP
)

// PWM
const (
	PWM_MODE_MS  = iota // C.PWM_MODE_MS
	PWM_MODE_BAL = iota // C.PWM_MODE_BAL
)

// Interrupt levels
const (
	INT_EDGE_SETUP   = iota // C.INT_EDGE_SETUP
	INT_EDGE_FALLING = iota // C.INT_EDGE_FALLING
	INT_EDGE_RISING  = iota // C.INT_EDGE_RISING
	INT_EDGE_BOTH    = iota // C.INT_EDGE_BOTH
)

// Conversion from wiringPi numeration to GPIO
func PinToGpio(pin int) int {
	return int(C.wpiPinToGpio(C.int(pin)))
}

// Core wiringPi functions

func WiringPiSetup() error {
	if -1 == int(C.wiringPiSetup()) {
		return errors.New("wiringPiSetup failed to call")
	}
	return nil
}

func WiringPiSetupSys() error {
	if -1 == int(C.wiringPiSetupSys()) {
		return errors.New("wiringPiSetupSys failed to call")
	}
	return nil
}

func WiringPiSetupGpio() error {
	if -1 == int(C.wiringPiSetupGpio()) {
		return errors.New("wiringPiSetupGpio failed to call")
	}
	return nil
}

func WiringPiSetupPhys() error {
	if -1 == int(C.wiringPiSetupPhys()) {
		return errors.New("wiringPiSetupPhys failed to call")
	}
	return nil
}

func PinModeAlt(pin int, mode int) {
	C.pinModeAlt(C.int(pin), C.int(mode))
}

func PinMode(pin int, mode int) {
	C.pinMode(C.int(pin), C.int(mode))
}

func PullUpDnControl(pin int, pud int) {
	C.pullUpDnControl(C.int(pin), C.int(pud))
}

func DigitalRead(pin int) int {
	return int(C.digitalRead(C.int(pin)))
}

func DigitalWrite(pin int, value int) {
	C.digitalWrite(C.int(pin), C.int(value))
}

func PwmWrite(pin int, value int) {
	C.pwmWrite(C.int(pin), C.int(value))
}

func AnalogRead(pin int) int {
	return int(C.analogRead(C.int(pin)))
}

func AnalogWrite(pin int, value int) {
	C.analogWrite(C.int(pin), C.int(value))
}

// On-Board Raspberry Pi hardware specific stuff

// Extras from arduino land

func Delay(ms uint) {
	C.delay(C.uint(ms))
}

func DelayMicroseconds(microSec uint) {
	C.delayMicroseconds(C.uint(microSec))
}
