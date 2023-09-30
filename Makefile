run:
	go run main.go
lint:
	golangci-lint run ./...
vuln:
	govulncheck ./...
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
