package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// concorrencia != PARALELISMO
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() // -1
	}()
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
