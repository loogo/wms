build:
	protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto

run:
	docker run --name demo_srv \
		-e MICRO_REGISTRY=consul \
		-e MICRO_REGISTRY_ADDRESS=consul:8500 \
		--link consul:consul \
		consignment-service