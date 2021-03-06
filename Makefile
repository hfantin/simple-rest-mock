all: clean update build-all

clean: 
	rm -rf bin

update: 
	go get -u
	go mod tidy

build-all: build-linux build-arm build-mac build-win

build-linux:
	GOOS=linux go build -o bin/srm

build-arm:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/srm-arm

build-mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/srm-mac

build-win:
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o bin/srm.exe

docker: 
	docker build -t srm .
