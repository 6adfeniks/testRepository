package main

/*
Первая часть
*/
//import (
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//)
//
//func Greet(writer io.Writer, name string) {
//	fmt.Fprintf(writer, "Hello, %s", name)
//}
//
//func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
//	Greet(w, "world")
//}
//
//func main() {
//	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
//}

/*
Вторая часть
*/

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(b io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(b, i)
	}
	sleeper.Sleep()
	fmt.Fprint(b, finalWord)

}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
