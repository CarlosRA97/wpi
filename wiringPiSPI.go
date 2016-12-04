package wpi

/*

#cgo LDFLAGS: -lwiringPi
#include <wiringPiSPI.h>
#include <stdlib.h>

*/
import "C"
import "unsafe"

type WiringPiSPI struct {
	channel int
}

func (wpiSPI *WiringPiSPI) GetFd() int {
	return int(C.wiringPiSPIGetFd(C.int(wpiSPI.channel)))
}

// TODO: Check this function pls, may have some problems with C pointers
func (wpiSPI *WiringPiSPI) DataRW(data string, len int) int {
	_data := *C.uchar(data)
	defer C.free(unsafe.Pointer(_data))
	return int(C.wiringPiSPIDataRW(C.int(wpiSPI.channel), _data, C.int(len)))
}

func (wpiSPI *WiringPiSPI) SetupMode(speed, mode int) int {
	return int(C.wiringPiSPISetupMode(C.int(wpiSPI.channel), C.int(speed), C.int(mode)))
}

func (wpiSPI *WiringPiSPI) Setup(speed int) int {
	return int(C.wiringPiSPISetup(C.int(wpiSPI.channel), C.int(speed)))
}
