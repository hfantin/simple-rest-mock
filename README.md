![GitHub release (latest by date)](https://img.shields.io/github/v/release/hfantin/simple-rest-mock)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hfantin/simple-rest-mock)
![GitHub repo size](https://img.shields.io/github/repo-size/hfantin/simple-rest-mock)
![GitHub](https://img.shields.io/github/license/hfantin/simple-rest-mock)


# SRM - Simple Rest Mock
This project aims to intercept the target server request and return a mock json to the chosen endpoints.

### Building this project:
- using [goreleaser](https://goreleaser.com/install/): 
> goreleaser release --snapshot --clean 
- using makefile: 
> make
  
### Downloading binaries: 
1. Choose a binary [here](https://github.com/hfantin/simple-rest-mock/releases) according to your operating system    
2. create a new .env file or set the env variables: 
```
SERVER_PORT=9000
TARGET_SERVER=<YOUR_TARGET_SERVER>
WRITE_FILE=true
USE_HTTPS=false
ENDPOINTS=/v1/endpoint1;v1/endpoint2;
```
3. start the server   
- linux and macos   
> chmod +x srm && ./srm
-  windows   
> srm.exe 

Obs.: at first, use the **WRITE_FILE=true** to record the responses in .files/ folder, then change it to false and restart the server when it is done.

### Errors 
- If you are receiving the message "cannot be opened because the developer cannot be verified" when executing on macos, try this command bellow:   
> xattr -d com.apple.quarantine simple-rest-mock
### Environment variables
```
SERVER_PORT: number of the where server runs, the default is 9000   
WRITE_FILE: this flag enables recording requests to a file   
TARGET_SERVER: this is target server where SCM will make request and record the response when WRITE_FILE is enabled
USE_HTTPS= to use https 
CERTIFICATE_PATH=certs/simple-rest-mock.crt
KEY_PATH=certs/simple-rest-mock.key
ENDPOINTS= list of endpoints intercepted separated by ;
```

### Rest api for tests
- [dogs](https://dog.ceo/api/breeds/image/random)

### using https: 
- create a certificate: 
> openssl genrsa -out simple-rest-mock.key 2048
- generate self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
> openssl req -new -x509 -sha256 -key simple-rest-mock.key -out simple-rest-mock.crt -days 3650
- for jvm access, import the self-signed certificate: 
> keytool -import -trustcacerts -keystore JVM_HOME/lib/security/cacerts -storepass changeit -noprompt -alias mycert -file simple-rest-mock.crt


### creating a release
> git tag -a v0.x.x -m "v0.x.x" && git push origin v0.x.x
> export GITHUB_TOKEN=<TOKEN_HERE>
- publish release 
> goreleaser release --skip-publish
- local release
> goreleaser release --snapshot --clean 

### links
- [How to Publish Your Golang Binaries with Goreleaser](https://www.kosli.com/blog/how-to-publish-your-golang-binaries-with-goreleaser/)
