# Go GRPC + Gateway + Buf Example

This repo is supposed to help with learning how to handle grpc/proto/grpc-gateway/openapi and more via a setup that can generate multiple different protobufs for different services inside your monorepo at the same time and allowing the owner to code using grpc and still expose a Json http 1.1 TCP transport layer if the client needs it.

## Requirements

GoExampleGRPC requires a golang to be installed correctly and it also requires protoc to be installed to work.
We have made this easy by running `make install_proto` to install the protoc required packages and then `make install_buf`
to install the buf package [more info](https://buf.build/docs/installation).

Buf will be responsible for generating the auto generated code (structs and interfaces) based on the [idl folder](./idl/proto/) files.
All the configurations for buf live in [gen file](./buf.gen.yaml) or the [work file](./buf.work.yaml).
This is where the output location for the generated files live and where it should look for input too + all the plugins that
you want to use in your code generation.

In this case we are generating `grpc`, `grpc-gateway` and `openapi specs` from the single proto location by using the specified plugins.
On important notice: [Google files located in the idl folder](./idl/proto/google/) were manually copied over and are required for grpc-gateway to work. This was available internally in the grpc golang package (in the third_party package), but now have to be downloaded and added manually.

## Usage

If you choose to generate the [idl_gen](./idl_gen/go/) `.pb.go` files - you need to run `make clean_proto` which will delete all the generated
files (golang, as an example) and then running `make proto` generates these based on buf.

After the golang files have been generated successfully, the next step is to simply go into your own service and add to the `go.mod` file, this:

```go
// Mapping to generated idl files (grpc)
replace github.com/copito/example_grpc/idl_gen/go => ../../idl_gen/go/
```

and run:

```bash
go mod tidy
```

The first part: `github.com/copito/example_grpc/idl_gen/go` should match the module name in the [idl_gen module](./idl_gen/go/go.mod)

```go
module github.com/copito/example_grpc/idl_gen/go
```

## Project

Your project directory should follow a specific structure:

```lua
idl/
|-- avro/
|-- proto/
| |-- google/
| |-- service1/
| | | --hello.proto
idl_gen/
|-- go/
| |-- service1/
| | | --hello.pb.go
| | | --hello_grpc.pb.go
| | | --hello.pb.gw.go
openapi/
services/
|-- service2/
|-- service1/
| |-- main.go
| |-- README.md
| |-- ...
-- README.md
-- Makefile
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the LICENSE [file](./LICENSE) for details.
Feel free to customize it further based on your project's specifics!
