module CodeMart

go 1.22.0

toolchain go1.22.2

require (
	github.com/go-sql-driver/mysql v1.9.2
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.4
	proto v0.0.0
)

replace proto => ./proto

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
