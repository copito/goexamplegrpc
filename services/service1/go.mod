module github.com/copito/service1

go 1.22.3

require (
	github.com/copito/example_grpc/idl_gen/go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.64.0
)

require google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240521202816-d264139d666e // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

// Mapping to generated idl files (grpc)
replace github.com/copito/example_grpc/idl_gen/go => ../../idl_gen/go/
