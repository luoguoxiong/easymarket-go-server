module easymarket-go-server

go 1.13

require (
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-kratos/kratos v0.4.2
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	github.com/google/wire v0.4.0
	github.com/jinzhu/gorm v1.9.12
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/urfave/cli/v2 v2.2.0 // indirect
	golang.org/x/sys v0.0.0-20200409092240-59c9f1ba88fa // indirect
	google.golang.org/genproto v0.0.0-20200409111301-baae70f3302d
	google.golang.org/grpc v1.28.1
)

replace google.golang.org/grpc v1.28.1 => google.golang.org/grpc v1.26.0
