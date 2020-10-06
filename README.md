
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/cmd ./cmd/main.go

kratos tool swagger serve app/api/api.swagger.json --port 3000