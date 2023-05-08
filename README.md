# SRM - Simple Rest Mock
This project is to help you mock responses from a rest api server.    

### How to build
- use **make** in the root folder, this will generate .bin folder with the binaries
  
### How to use
1. Choose a binary according to your OS in the ./bin folder:    
   a. srm for linux   
   b. srm.exe for windows   
   c. srm.app for macos   
2. Write your json files in the .files directory, like the example below:   
```
endpoint: /v1/items
file: .files/items.GET.json
```
3. create .env file: 
```
SERVER_PORT=9443
TARGET_SERVER=<YOUR_TARGET_SERVER>
WRITE_FILE=false
USE_HTTPS=false
CERTIFICATE_PATH=certs/simple-rest-mock.crt
KEY_PATH=certs/simple-rest-mock.key
ENDPOINTS=/v1/endpoint1;v1/endpoint2;
```
4. When you call for localhost:9000/v1/items, you will receive the content of the file as response.    


### Environment variables
SERVER_PORT: number of the where server runs, the default is 9000   
WRITE_FILE: this flag enables recording requests to a file   
TARGET_SERVER: this is target server where SCM will make request and record the response when WRITE_FILE is enabled
USE_HTTPS= to use https 
CERTIFICATE_PATH=certs/simple-rest-mock.crt
KEY_PATH=certs/simple-rest-mock.key
ENDPOINTS= list of endpoints separated by ;

### Rest api for tests
- [dogs](https://dog.ceo/api/breeds/image/random)

### using https: 
- create a certificate: 
> openssl genrsa -out simple-rest-mock.key 2048
- generate self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
> openssl req -new -x509 -sha256 -key simple-rest-mock.key -out simple-rest-mock.crt -days 3650
- for jvm access, import the self-signed certificate: 
> keytool -import -trustcacerts -keystore JVM_HOME/lib/security/cacerts -storepass changeit -noprompt -alias mycert -file simple-rest-mock.crt
