build:
	protoc --proto_path=$(GOPATH)/src:. --go_out=. --micro_out=. proto/consignment/consignment.proto
	# GOOS=linux GOARCH=amd64 go build
	docker build --build-arg https_proxy=http://192.168.0.100:1087 -t consignment-service .

run: 
	docker run --name demo_srv \
		-e MICRO_REGISTRY=consul \
		-e MICRO_REGISTRY_ADDRESS=consul:8500 \
		--link consul:consul \
		consignment-service