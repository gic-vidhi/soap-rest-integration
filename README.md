# This service is used to connect with soap api's from rest api's

## Define port number in environment variable 
```
variable name : SOAP_REST_PORT
port number : 9002
```

## Run service
```
go run main.go
```

#### Call a soap get endpoint 

```http
  GET localhost:9002/soap-rest-integration-service/get-from-soap-api/?country=US
```
```curl
curl --location 'localhost:9002/soap-rest-integration-service/get-from-soap-api/?country=US'
```

#### Call a soap post endpoint
```http
  POST localhost:9002/soap-rest-integration-service/post-to-soap-api/?country=US
```

```curl
curl --location --request POST 'localhost:9002/soap-rest-integration-service/post-to-soap-api/?country=US'
```

## Authors

- [Vidhi Goel](https://www.github.com/gic-vidhi)