proto:
		protoc pkg/pb/*.proto --go-grpc_out=. --go_out=. 
		protoc pkg/pb/userproto/*.proto --go-grpc_out=. --go_out=. 
server:
	go run cmd/main.go	