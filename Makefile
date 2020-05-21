all: clean update compile docker

clean: 
	rm -rf bin

update: 
	go get -u
	go mod tidy

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/srm-arm64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/srm-mac
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o bin/srm.exe
	GOOS=linux go build -o bin/srm

docker: 
	docker build -t srm .
