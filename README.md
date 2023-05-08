# SRM - Simple Rest Mock
This project aims to intercept your requests to a target server and return a json mock to the chosen endpoints.

### How to build
- use **make** in the root folder, this will generate .bin folder with the binaries
  
### How to use
1. Choose a binary [here](https://github.com/hfantin/simple-rest-mock/releases)    
2. unzip your binary
3. create a new .env file or the env variables: 
```
SERVER_PORT=9443
TARGET_SERVER=<YOUR_TARGET_SERVER>
WRITE_FILE=false
USE_HTTPS=false
ENDPOINTS=/v1/endpoint1;v1/endpoint2;
```
4. start the server   
- linux and macos   
> chmod +x simple-rest-mock &&./simple-rest-mock 
-  windows   
> simple-rest-mock.exe 

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
> goreleaser release --skip-publish 

### links
- [How to Publish Your Golang Binaries with Goreleaser](https://www.kosli.com/blog/how-to-publish-your-golang-binaries-with-goreleaser/)