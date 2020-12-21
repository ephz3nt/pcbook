gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=pb --go_out=pb
clean:
	rm pb/*.go
fmt:
	go fmt ./...
run:fmt
	go run main.go
test:fmt
	go test -cover -race ./...
.PHONY: gen clean fmt run test
