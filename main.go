package main

import "sync"

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		startHttp()
	}()

	go func() {
		defer wg.Done()
		startGrpc()
	}()

	wg.Wait()
}
