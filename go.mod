module github.com/developerc/finprojagent3

go 1.22.1

replace grpc => ./grpc/

require (
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	grpc v0.0.0-00010101000000-000000000000
)

require (
	github.com/dengsgo/math-engine v0.0.0-20230823154425-78f211b48149 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
)
