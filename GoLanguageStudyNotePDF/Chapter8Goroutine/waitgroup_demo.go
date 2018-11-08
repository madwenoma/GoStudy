package main

import (
	"sync"
	"time"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	for i := 0; i < 10; i++ {
		wg.Add(1) //要放在go fun外面，否则有可能来不及设置

		go func(x int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("goroute:", x, "done")
		}(i)
	}

	fmt.Println("main...")
	wg.Wait()

	fmt.Println("main done...")
	fmt.Println("===========test moreThanOneWait============")
	moreThanOneWait()
}

func moreThanOneWait() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()
		fmt.Println("put the elephant")
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("open the door")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("close the door")
}
