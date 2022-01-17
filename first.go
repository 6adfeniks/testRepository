package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0;i<10;i++ {
		go Factorial(i)
	}
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000  * time.Millisecond)
	go func() {
		for {
			select {
			case <-tick:
				fmt.Println("tick")
			case <-boom:
				fmt.Println("Boom")
				return
			}
		}
	}()
	time.Sleep(3*time.Second)

}

func Factorial(i int) int{
	switch i{
	case 0:
		fmt.Println(1)
		return 1
	default:
		res := 1
		for n:=1;n<=i;n++{
			res *= n
		}
		fmt.Println(res)
		return res
	}
}
