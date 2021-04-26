package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep(time.Duration)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep(sleepTime time.Duration) {
	time.Sleep(sleepTime)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
