GOPATH:=$(shell go env GOPATH)

PHONY: install-deps
install-deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/google/wire/cmd/wire@latest

PHONY: migrate
migrate:
	goose -dir=migrations postgres "postgres://postgres:secret@localhost:5432/users" up

PHONY: compose
compose:
	docker-compose -f ./build/docker-compose.yaml up -d

PHONY: injection
injection:
	wire ./internal/injection

PHONY: gen
gen:
	gen-user-api

PHONY: gen-user-api
gen-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path api \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	api/user_v1/user.proto