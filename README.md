![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hfantin/simple-rest-mock)
![GitHub](https://img.shields.io/github/license/hfantin/simple-rest-mock)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/hfantin/simple-rest-mock)
![GitHub repo size](https://img.shields.io/github/repo-size/hfantin/simple-rest-mock)

# SRM - Simple Rest Mock
Simple Rest Mock is a request/response interceptor that can replace the response of the target server returning the mock file content

```
Usage:
  srm [flags]

Flags:
      --config string                config file (default is $HOME/.srm/config.yaml)
  -e, --endpoints strings            endpoints filtered by regex
  -h, --help                         help for srm
  -p, --port string                  server port (default "9000")
  -r, --rec-mode                     recorde response
  -f, --response-files-path string   path to write response files (default "jsons")
  -t, --target-server string         target server to intercept request/response
  -v, --version                      version for srm
```

## Download and run: 
1. Choose a binary [here](https://github.com/hfantin/simple-rest-mock/releases) according to your operating system and rename it as **srm**
2. use the flags or create a new $HOME/.srm/config.yaml file with the content below: 
```
port: 5000                       # default is 9000
rec-mode: true                   # default is false
response-files-path: resp-jsons  # default is $HOME/.srm/jsons
target-server: https://catfact.ninja
endpoints:
  - /breeds
  - /fact
  - /facts
```
3. start the server   
- linux and macos: `cmd +x srm && ./srm`   
-  windows: `srm.exe`   

Obs.: when **rec-response** is enabled, every request to the endpoints in the list will be intercepted and its responses recorded into $HOME/.srm/<response-files-path> folder. 

### Errors 
- If you are receiving the message "cannot be opened because the developer cannot be verified" when executing on macos, try this command: `xattr -d com.apple.quarantine srm`   

### Rest api for tests
- [dogs](https://dog.ceo/api/breeds/image/random)
- [cats breeds](https://catfact.ninja/breeds) 
- [cats fact](https://catfact.ninja/facts) 
- [cats facts](https://catfact.ninja/fact) 

## Development

### Building this project:
- using [goreleaser](https://goreleaser.com/install/): 
> goreleaser release --snapshot --clean 
- using makefile: 
> make

### creating a release
> git tag -a v0.x.x -m "v0.x.x" && git push origin v0.x.x
> export GITHUB_TOKEN=<TOKEN_HERE>
- publish release 
> goreleaser release --skip-publish
- local release
> goreleaser release --snapshot --clean 

### using https: 
- create a certificate: 
> openssl genrsa -out simple-rest-mock.key 2048
- generate self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
> openssl req -new -x509 -sha256 -key simple-rest-mock.key -out simple-rest-mock.crt -days 3650
- for jvm access, import the self-signed certificate: 
> keytool -import -trustcacerts -keystore JVM_HOME/lib/security/cacerts -storepass changeit -noprompt -alias mycert -file simple-rest-mock.crt

### links
- [How to Publish Your Golang Binaries with Goreleaser](https://www.kosli.com/blog/how-to-publish-your-golang-binaries-with-goreleaser/)
