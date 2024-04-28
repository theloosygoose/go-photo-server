run:
	@go run cmd/photo-server/main.go

build:
	@go build cmd/photo-server/main.go
	@mv ./main ./bin/main
