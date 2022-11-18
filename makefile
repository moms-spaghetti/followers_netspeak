proto:
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/followers.proto

build:
		go build cmd/main.go

run:
		./main

build_run:
		go build cmd/main.go && ./main