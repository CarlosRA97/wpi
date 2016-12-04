package wpi

/*

#cgo LDFLAGS: -lwiringPi
#include <wiringPiI2C.h>
#include <stdlib.h>

*/
import "C"
import "unsafe"

type I2C int

func I2CSetupInterface(device string, devID int) I2C {
	var _new I2C
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	_new = I2C(C.wiringPiI2CSetupInterface(dev, C.int(devID)))
	return _new
}

func I2CSetup(devID int) I2C {
	return I2C(C.wiringPiI2CSetup(C.int(devID)))
}

func (i2c I2C) Read() int {
	return int(C.wiringPiI2CRead(C.int(i2c)))
}
func (i2c I2C) ReadReg8(reg int) int {
	return int(C.wiringPiI2CReadReg8(C.int(i2c), C.int(reg)))
}
func (i2c I2C) ReadReg16(reg int) int {
	return int(C.wiringPiI2CReadReg16(C.int(i2c), C.int(reg)))
}

func (i2c I2C) Write(data int) int {
	return int(C.wiringPiI2CWrite(C.int(i2c), C.int(data)))
}
func (i2c I2C) WriteReg8(reg, data int) int {
	return int(C.wiringPiI2CWriteReg8(C.int(i2c), C.int(reg), C.int(data)))
}
func (i2c I2C) WriteReg16(reg, data int) int {
	return int(C.wiringPiI2CWriteReg16(C.int(i2c), C.int(reg), C.int(data)))
}
