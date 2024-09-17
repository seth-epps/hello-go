package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

func startHttp() error {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGet).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
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

func handleGet(w http.ResponseWriter, req *http.Request) {
	res := Response{IP: req.RemoteAddr, Message: "Hello From Go!"}
	json.NewEncoder(w).Encode(res)
}
