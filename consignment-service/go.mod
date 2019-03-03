module github.com/loogo/wms/consignment-service

go 1.12

require (
	github.com/golang/protobuf v1.3.0
	github.com/loogo/wms/vessel-service v0.0.0
	github.com/micro/go-micro v0.27.0
)

replace github.com/loogo/wms/vessel-service v0.0.0 => ../vessel-service
