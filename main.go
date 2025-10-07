package main

import (
	"flag"
	"sync"
)

const defaultServerName = "Go"

func main() {

	serverOpts := serverOpts{}

	flag.StringVar(&serverOpts.name, "server-name", defaultServerName, "server name to report back to client")
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		startHttp(serverOpts)
	}()

	go func() {
		defer wg.Done()
		startGrpc(serverOpts)
	}()

	wg.Wait()
}
