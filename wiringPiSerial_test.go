package wpi

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	serial1, err := Open("/dev/tty00", 9600)
	defer serial1.Close()
	if err != nil {
		fmt.Println("Something wrong happened. 404")
	}

}
