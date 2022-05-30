module simple-douyin

go 1.17

require (
	github.com/cloudwego/kitex v0.3.1
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/kitex-contrib/registry-etcd v0.0.0-20220110034026-b1c94979cea3
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/tidwall/gjson v1.12.1 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
	gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
