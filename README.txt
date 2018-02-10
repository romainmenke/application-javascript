- go run main.go
- visit : localhost:8000
- content-type : application/javascript

- docker build -t application-javascript .
- docker run -it -p 8000:8000 application-javascript
- visit : localhost:8000
- content-type : application/x-javascript

- change fixXJavascript flag to true
- docker build -t application-javascript .
- docker run -it -p 8000:8000 application-javascript
- visit : localhost:8000
- content-type : application/javascript
