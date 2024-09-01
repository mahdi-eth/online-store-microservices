PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GRPC_GO = protoc-gen-go-grpc

.PHONY: proto

proto:
	@for service in product order inventory; do \
 		protoc --proto_path=. --go_out=./services/$$service/grpc \
        --go_opt=paths=source_relative proto/api/$$service/*.proto; \
	done

docker-proto:
	docker build -t grpc-generator . && docker run --rm -v $(pwd):/app grpc-generator

docker:
	docker-compose up --build -d

k8s:
	