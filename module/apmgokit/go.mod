module go.elastic.co/apm/module/apmgokit

require (
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/stretchr/testify v1.8.3
	go.elastic.co/apm v1.11.0
	go.elastic.co/apm/module/apmgrpc v1.11.0
	go.elastic.co/apm/module/apmhttp v1.11.0
	golang.org/x/net v0.16.0
	google.golang.org/grpc v1.58.2
	google.golang.org/grpc/examples v0.0.0-20231024221354-6e14274d008d
)

replace go.elastic.co/apm => ../..

replace go.elastic.co/apm/module/apmgrpc => ../apmgrpc

replace go.elastic.co/apm/module/apmhttp => ../apmhttp

go 1.13
