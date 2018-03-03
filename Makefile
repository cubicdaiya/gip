all: gip

gip:
	go build

fmt:
	go fmt ./...

clean:
	rm -rf gip
