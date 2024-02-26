generate-proto:
	protoc \
    -I=internal/api/grpc/proto \
    --go_opt=paths=source_relative \
    --go_out=internal/api/grpc/proto_gen \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=internal/api/grpc/proto_gen \
	--go-grpc_opt=require_unimplemented_servers=false \
	internal/api/grpc/proto/*.proto

generate-proto-kafka:
	protoc \
    -I=internal/api/kafka/proto \
    --go_opt=paths=source_relative \
    --go_out=internal/api/kafka/proto_gen \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=internal/api/kafka/proto_gen \
	--go-grpc_opt=require_unimplemented_servers=false \
	internal/api/kafka/proto/*.proto

migrate-up:
ifdef step
	./apollo_eks_base_service migrate up --step=${step}
else
	./apollo_eks_base_service migrate up
endif

migrate-down:
	./apollo_eks_base_service migrate down --step=${step}

create-migration-files:
	migrate create -ext sql -dir migrations ${filename}

swagger:
	swag init

install:
	sh install.sh
