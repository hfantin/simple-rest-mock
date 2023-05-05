# SRM - Simple Rest Mock
This project is to help you mock responses from a rest api server.    

### How to use
1. Build the project using **make** in the root.
2. Choose a binary, for example, **./bin/srm** 
3. Write your json files in the .files directory, like the example below:   
```
endpoint: /v1/items
file: .files/items.GET.json
```
4. create .env file: 
```
SERVER_PORT=9443
TARGET_SERVER=<YOUR_TARGET_SERVER>
WRITE_FILE=true
CERTIFICATE_PATH=certs/simple-rest-mock.crt
KEY_PATH=certs/simple-rest-mock.key
```
5. When you call for localhost:5000/v1/items, you will receive the content of the file as response.    

### Environment variables
SERVER_PORT: number of the where server runs, the default is 5000   
WRITE_FILE: this flag enables recording requests to a file   
TARGET_SERVER: this is target server where SCM will make request and record the response when WRITE_FILE is enabled

### Rest api for tests
- [dogs](https://dog.ceo/api/breeds/image/random)

### using https: 
- create a certificate: 
> openssl genrsa -out simple-rest-mock.key 2048
- generate self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
> openssl req -new -x509 -sha256 -key simple-rest-mock.key -out simple-rest-mock.crt -days 3650
- for jvm access, import the self-signed certificate: 
> keytool -import -trustcacerts -keystore JVM_HOME/lib/security/cacerts -storepass changeit -noprompt -alias mycert -file simple-rest-mock.crt