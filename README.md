# SRM - Simple Rest Mock
This project is to help you to mock the responses from a rest api server.    

### How to use
1. Build the project using **make** in the root.
2. Choose a binary, for example, **./bin/srm** 
3. Write your json files in the .files directory, like the example below:   
```
endpoint: /v1/hello
file: .files/v1.hello.json
```
4. When you call for localhost:5000/v1/hello, you will receive the content of the file as response.


### Rest api for tests
- [dogs](https://dog.ceo/api/breeds/image/random)

