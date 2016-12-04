package wpi

/*

#cgo LDFLAGS: -lwiringPi
#include <wiringSerial.h>
#include <stdlib.h>

void _serialPrintf(int fd, char* message) {
	serialPrintf(fd, message);
}

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Serial int

const _nil = -1

func Open(device string, baud int) (*Serial, error) {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	serial := *Serial(C.serialOpen(dev, C.int(baud)))
	if serial != _nil {
		return *Serial(C.serialOpen(dev, C.int(baud))), nil
	} else {
		return nil, errors.New(fmt.Sprintf("Unable to open serial device: %s", device))
	}
}
func (s *Serial) Close() {
	C.serialClose(C.int(s))
}
func (s *Serial) Flush() {
	C.serialFlush(C.int(s))
}
func (s *Serial) Putchar(c string) {
	_c := C.CString(c)
	defer C.free(unsafe.Pointer(_c))
	C.serialPutchar(C.int(s), _c)
}
func (s *Serial) Puts(ss string) {
	_s := C.CString(ss)
	defer C.free(unsafe.Pointer(_s))
	C.serialPuts(C.int(s))
}
func (s *Serial) Printf(message string) {
	_message := C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	C._serialPrintf(C.int(s), _message)
}
func (s *Serial) DataAvail() int {
	return int(C.serialDataAvail(C.int(s)))
}
func (s *Serial) Getchar() int {
	return int(C.serialGetchar(C.int(s)))
}
