package vunikbd

import (
	"fmt"
	"time"

	"github.com/bendahl/uinput"
)

type Keyboard struct {
	name  string
	delay time.Duration
	vk    uinput.Keyboard
}

func NewKeyboard(name string, delay time.Duration) (*Keyboard, error) {
	var err error

	k := Keyboard{
		name:  name,
		delay: delay,
	}

	k.vk, err = uinput.CreateKeyboard("/dev/uinput", k.name)

	return &k, err
}

func (k *Keyboard) Type(s string) {
	for _, r := range s {
		h := fmt.Sprintf("%x\n", r)
		k.vk.SendKeyPress(uinput.KEY_LEFTCTRL)
		k.sleep()
		k.vk.SendKeyPress(uinput.KEY_LEFTSHIFT)
		k.sleep()
		k.vk.SendKeyPress(uinput.KEY_U)
		k.sleep()
		k.vk.SendKeyRelease(uinput.KEY_U)
		k.sleep()
		for _, v := range h {
			k.vk.SendKeyPress(keycodes[v])
			k.sleep()
			k.vk.SendKeyRelease(keycodes[v])
			k.sleep()
		}
		k.vk.SendKeyRelease(uinput.KEY_LEFTCTRL)
		k.sleep()
		k.vk.SendKeyRelease(uinput.KEY_LEFTSHIFT)
		k.sleep()
	}
}

func (k *Keyboard) sleep() {
	time.Sleep(k.delay)
}

var keycodes map[rune]int

func init() {
	keycodes = make(map[rune]int)
	keycodes['0'] = uinput.KEY_KP0
	keycodes['1'] = uinput.KEY_KP1
	keycodes['2'] = uinput.KEY_KP2
	keycodes['3'] = uinput.KEY_KP3
	keycodes['4'] = uinput.KEY_KP4
	keycodes['5'] = uinput.KEY_KP5
	keycodes['6'] = uinput.KEY_KP6
	keycodes['7'] = uinput.KEY_KP7
	keycodes['8'] = uinput.KEY_KP8
	keycodes['9'] = uinput.KEY_KP9
	keycodes['a'] = uinput.KEY_A
	keycodes['b'] = uinput.KEY_B
	keycodes['c'] = uinput.KEY_C
	keycodes['d'] = uinput.KEY_D
	keycodes['e'] = uinput.KEY_E
	keycodes['f'] = uinput.KEY_F
}
