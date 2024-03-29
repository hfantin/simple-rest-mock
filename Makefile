VERSION:=$(shell sed -nE '/version/{s/.*:\s*"(.*)",/\1/p;q;}' build.json | xargs)
LD_FLAGS:=-ldflags "-X 'github.com/hfantin/simple-rest-mock/config.versionNumber=${VERSION}'"

all: print-version clean update build-all

release: 
	@git tag -af v${VERSION} -m "v${VERSION}"
	@git push origin v${VERSION}
	@goreleaser release
	
print-version: 
	@echo "building version ${VERSION}"

clean: 
	@rm -rf dist

update: 
	@go get -u
	@go mod tidy

build-all: build-linux build-mac build-win

build-linux:
	@GOOS=linux go build ${LD_FLAGS} -o dist/srm

build-arm:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ${LD_FLAGS} -o dist/srm.arm

build-mac:
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build ${LD_FLAGS} -o dist/srm.app

build-win:
	@GOOS=windows GOARCH=386 CGO_ENABLED=0 go build ${LD_FLAGS} -o dist/srm.exe

copy-certificates: 
	@cp certs/* ./dist

docker: 
	@docker build -t srm .
