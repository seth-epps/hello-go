# hello-go
Simple json and grpc web server written in go with [gorilla/mux](https://github.com/gorilla/mux)

This is _not_ intended to be used in any meaningful way other than to test in simple local and container runtimes.


## Usage
If building from scratch always run `make generate` before hand.
### Docker
```sh
docker image build -t hello-go .
docker container run --rm -d -p 8080:8080 -p 9090:9090 --name hello-go hello-go
curl localhost:8080/
# {"ip":"192.168.65.1:61407","message":"Hello From Go!"}
grpcurl -plaintext localhost:9090 main.Hello/SayHello
# {
#   "ip": "192.168.65.1:58783",
#   "message": "Hello From Go!"
# }
```

### Locally
```sh
go run .
# json http server listening at [::]:8080
# gRPC server listening at [::]:9090
curl localhost:8080/
# {"ip":"[::1]:53963","message":"Hello From Go!"}
grpcurl -plaintext localhost:9090 main.Hello/SayHello
# {
#   "ip": "[::1]:53973",
#   "message": "Hello From Go!"
# }
```