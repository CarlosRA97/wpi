package wpi

/*

#cgo LDFLAGS: -lwiringPi
#include <wiringSerial.h>
#include <stdlib.h>

*/
import "C"
import "unsafe"

type Serial int

func Open(device string, baud int) *Serial {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	return *int(C.serialOpen(dev, C.int(baud)))
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
	C.serialPrintf(_message)
}
func (s *Serial) DataAvail() int {
	return int(C.serialDataAvail(C.int(s)))
}
func (s *Serial) Getchar() int {
	return int(C.serialGetchar(C.int(s)))
}
