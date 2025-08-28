package throbber

import (
	"fmt"
	"time"
	"os"
)

const chars uint32 = ('/') | ('-' << 0x08) | ('\\' << 0x10) | ('|' << 0x18)

type throbber struct {
	run bool
	step int
	delay time.Duration
	wait_til time.Time
}

var throb throbber = throbber{
	run: false,
	delay: time.Millisecond * 250,
	wait_til: time.Time{},
}

func Start() {
	throb.run = true
	go func() {
		for throb.run {
			if throb.wait_til.After(time.Now()) {
				time.Sleep(throb.wait_til.Sub(time.Now()))
			}
			if !throb.run {
				return
			}
			char := uint8(0xff & (chars >> (throb.step * 8)))
			fmt.Fprintf(os.Stderr, "%c\r", char)
			throb.step = (throb.step + 1) & 3
			time.Sleep(throb.delay)
		}
	}()
}

func Stop() {
	throb.run = false
}

func Delay() {
	throb.wait_til = time.Now().Add(time.Millisecond * 300)
}
