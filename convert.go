package wpi

import "fmt"

// Conversions from GPIO and Board to wiringPi numeration
var (
	board2pin = []int{
		-1,
		-1,
		-1,
		PIN_SDA,
		-1,
		PIN_SCL,
		-1,
		PIN_GPIO_7,
		PIN_TXD,
		-1,
		PIN_RXD,
		PIN_GPIO_0,
		PIN_GPIO_1,
		PIN_GPIO_2,
		-1,
		PIN_GPIO_3,
		PIN_GPIO_4,
		-1,
		PIN_GPIO_5,
		PIN_MOSI,
		-1,
		PIN_MISO,
		PIN_GPIO_6,
		PIN_SCLK,
		PIN_CE0,
		-1,
		PIN_CE1,
	}
	gpio2pin = []int{
		PIN_SDA,
		PIN_SCL,
		-1,
		-1,
		PIN_GPIO_7,
		-1,
		-1,
		PIN_CE1,
		PIN_CE0,
		PIN_MISO,
		PIN_MOSI,
		PIN_SCLK,
		-1,
		-1,
		PIN_TXD,
		PIN_RXD,
		-1,
		PIN_GPIO_0,
		PIN_GPIO_1,
		-1,
		-1,
		PIN_GPIO_2,
		PIN_GPIO_3,
		PIN_GPIO_4,
		PIN_GPIO_5,
		PIN_GPIO_6,
		-1,
		-1,
		PIN_GPIO_8,
		PIN_GPIO_9,
		PIN_GPIO_10,
		PIN_GPIO_11,
	}
)

//use RPi.GPIO's BOARD numbering
func BoardToPin(pin int) int {
	if pin < 1 || pin >= len(board2pin) {
		panic(fmt.Sprintf("Invalid board pin number: %d", pin))
	}
	return board2pin[pin]
}

func GpioToPin(pin int) int {
	if pin < 0 || pin >= len(gpio2pin) {
		panic(fmt.Sprintf("Invalid bcm gpio number: %d", pin))
	}
	return gpio2pin[pin]
}
