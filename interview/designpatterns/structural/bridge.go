package structural

import "fmt"

type Device interface {
	On()
	Off()
}

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV is ON")
}

func (t *TV) Off() {
	fmt.Println("TV is OFF")
}

type Remote interface {
	Power()
}

type BasicRemote struct {
	device Device
}

func (r *BasicRemote) Power() {
	fmt.Println("Remote: power toggle")
	if r.device != nil {
		r.device.On()
		r.device.Off()
	}
}

func Bridge() {
	tv := &TV{}
	remote := &BasicRemote{device: tv}

	remote.Power()
}
