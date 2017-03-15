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

func (k *Keyboard) TypeRune(r rune) {
	h := fmt.Sprintf("%x\n", r)

	k.vk.SendKeyPress(uinput.KEY_LEFTCTRL)
	k.vk.SendKeyPress(uinput.KEY_LEFTSHIFT)
	k.vk.SendKeyPress(uinput.KEY_U)
	k.vk.SendKeyRelease(uinput.KEY_U)

	for _, v := range h {
		k.vk.SendKeyPress(keycodes[v])
		k.vk.SendKeyRelease(keycodes[v])
	}

	k.vk.SendKeyRelease(uinput.KEY_LEFTCTRL)
	k.vk.SendKeyRelease(uinput.KEY_LEFTSHIFT)
}

func (k *Keyboard) TypeString(s string) {
	for _, r := range s {
		k.TypeRune(r)
		time.Sleep(k.delay)
	}
}

func (k *Keyboard) KeyPress(code int) {
	k.vk.SendKeyPress(code)
	k.vk.SendKeyRelease(code)
}

var keycodes map[rune]int

func init() {
	keycodes = make(map[rune]int)
	keycodes['0'] = uinput.KEY_0
	keycodes['1'] = uinput.KEY_1
	keycodes['2'] = uinput.KEY_2
	keycodes['3'] = uinput.KEY_3
	keycodes['4'] = uinput.KEY_4
	keycodes['5'] = uinput.KEY_5
	keycodes['6'] = uinput.KEY_6
	keycodes['7'] = uinput.KEY_7
	keycodes['8'] = uinput.KEY_8
	keycodes['9'] = uinput.KEY_9
	keycodes['a'] = uinput.KEY_A
	keycodes['b'] = uinput.KEY_B
	keycodes['c'] = uinput.KEY_C
	keycodes['d'] = uinput.KEY_D
	keycodes['e'] = uinput.KEY_E
	keycodes['f'] = uinput.KEY_F
}
