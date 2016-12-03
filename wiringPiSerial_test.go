package wpi

import "testing"

func TestOpen(t *testing.T) {
	serial1 := Open("/dev/tty00", 9600)
	defer serial1.Close()
}
