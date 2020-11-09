# go-gin-starter-kit
The fast way to create a restful apis with Gin Framework with a structured project that defaults to mongodb and redis 

### how to start

copy & rename config/default-sample.json to default.json

### run in development mode
```
gin -i --notifications run start
```

### create swagger 
```
swag init
```
 http://localhost:8080/swagger/index.html

## CLI

### create controller
```
go run main.go generate controller {$name}     
```