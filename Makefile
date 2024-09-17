# Adding some helper makefile targets to generate gRPC stubs
.PHONY: generate
generate: tools
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protos/hello.proto


.PHONE: tools
tools:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest