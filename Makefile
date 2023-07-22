build:
	CGO_ENABLED=1 go build -o prizepicks-assessment ./main.go

run:
	CGO_ENABLED=1 go run ./main.go

test:
	CGO_ENABLED=1 go test -v -count=1 -race ./...

vendor:
	go mod vendor

tidy:
	go mod tidy

vet:
	go vet main.go

fmt:
	gofmt -d ./*.go ./models/*.go

staticcheck:
	staticcheck

golangcilint:
	golangci-lint run --no-config --enable-all --disable wrapcheck --disable varnamelen --disable paralleltest --disable nonamedreturns --disable nlreturn --disable tagliatelle --disable testpackage --disable depguard --disable wsl --disable lll --disable goerr113 --disable dupl --disable rowserrcheck --disable gci
