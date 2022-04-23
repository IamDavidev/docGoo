package main

import (
	"fmt"
	"time"
)

func main() {
	// concurencia forma 1 async
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	ContarObj("oveja")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// concurencia forma 2 con canal
	canal := make(chan string)
	go ContarObjChan("oveja", canal)

	for msj := range canal {
		fmt.Println("Objeto: ", msj)

	}

}

func ContarObj(obj string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Objeto: ", obj, " - ", i)
		time.Sleep(time.Second)
	}
}

func ContarObjChan(obj string, canal chan string) {

	for i := 0; i < 10; i++ {
		canal <- obj
		time.Sleep(time.Second)
	}
	close(canal)
}
