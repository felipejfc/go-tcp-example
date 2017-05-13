default: build

build:
	@mkdir -p bin && go build -o bin/gotcp main.go

build-linux:
	@mkdir -p bin && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/gotcp-linux-amd64 main.go

docker: build-linux
	@docker build -t gotcp .
