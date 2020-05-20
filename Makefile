all: clean update compile docker

clean: 
	rm -rf bin

update: 
	go get -u
	go mod tidy

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/simple-rest-mock-arm64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/simple-rest-mock-mac
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o bin/simple-rest-mock.exe
	GOOS=linux go build -o bin/simple-rest-mock

docker: 
	docker build -t simple-rest-mock .
