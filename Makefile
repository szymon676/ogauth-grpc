proto: 
	@echo making proto
	@protoc --go_out=pb-gen --go_opt=paths=source_relative --go-grpc_out=pb-gen --go-grpc_opt=paths=source_relative auth.proto
	@echo proto made succesfully	
run:
	@go run main.go