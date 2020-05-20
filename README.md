# Simple Rest Mock
This project is to help you to mock te response from server when you don't have access to server or the server is down.    

### How to use
1. Build the project using **make** in the root.
2. Choose a binary, for example, **./bin/simple-rest-mock** 
3. Write your json files in the .files directory, like the example below:   
```
endpoint: /v1/hello
file: .files/v1.hello.json
```
4. When you call for localhost:5000/v1/hello, you will as response the content of the file.

