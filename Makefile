# migrations
.PHONY: db
db:
	@docker-compose up -d postgres

.PHONY: migrate-create
migrate-create:
	@go run main.go migrate create $(name)

.PHONY: migrate-up
migrate-up: db
	@go run main.go migrate up

.PHONY: migrate-down
migrate-down: db
	@go run main.go migrate down

.PHONY: migrate-rollback
migrate-rollback: db
	@go run main.go migrate rollback

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-ci
test-ci:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: data
data:
	@go run main.go data

# protos
PROTOBUF_DIR = protobuf
PROTO_FILES = $(PROTOBUF_DIR)/*.proto
PROTOS_GO_OUT_PATH = pb
JS_OUT_DIR = client/pb

#proto
.PHONY: proto
proto:
	@mkdir -p pb client/pb
	@protoc --proto_path=${PROTOBUF_DIR} \
			 --go_out=plugins=grpc:$(PROTOS_GO_OUT_PATH) \
			 --js_out=import_style=commonjs:${JS_OUT_DIR} \
             --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${JS_OUT_DIR} \
			 ${PROTO_FILES}

.PHONY: server-grpc
server-grpc:
	@go run main.go server grpc
