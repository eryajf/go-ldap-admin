default: build

run:
	GIN_MODE=release go run main.go

build:
	go build -o go-ldap-admin main.go

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o go-ldap-admin main.go

build-linux-arm:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o go-ldap-admin main.go

lint:
	env GOGC=25 golangci-lint run --fix -j 8 -v ./... --timeout=5m --skip-files="public/client/feishu/feishu.go"