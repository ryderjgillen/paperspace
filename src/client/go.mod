module portService/client

// replace for non-VCS modules: https://github.com/golang/go/wiki/Modules#can-i-work-entirely-outside-of-vcs-on-my-local-filesystem
replace portService/grpc => ../grpc

go 1.18

require (
	google.golang.org/grpc v1.45.0
	portService/grpc v0.0.0-00010101000000-000000000000
)

require github.com/pkg/errors v0.9.1 // indirect

require (
	github.com/alecthomas/kong v0.5.0
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
