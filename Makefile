all: gip

gip: *.go
	go build

fmt:
	go fmt ./...

clean:
	rm -rf gip
