package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type Response struct {
	IP      string `json:"ip"`
	Host    string `json:"host"`
	Message string `json:"message"`
}

func startHttp(serverOpts serverOpts) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleGet(serverOpts.name))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// using a dedicated listener instead of ListenAndServe
	// to align with the grpc server
	//
	// effectively a stripped down implementation of the same thing
	lis, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen on port %v: %w", srv.Addr, err)
	}

	fmt.Printf("json http server listening at %v\n", lis.Addr())
	return srv.Serve(lis)
}

func handleGet(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		res := Response{
			IP:      req.RemoteAddr,
			Message: "Hello From " + name + "!",
			Host:    req.Host,
		}
		json.NewEncoder(w).Encode(res)
	}
}
