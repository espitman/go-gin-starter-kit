# go-gin-starter-kit
The fast way to create a restful apis with Gin Framework with a structured project that defaults to mongodb and redis 

### how to start

copy & rename config/default-sample.json to default.json

### run in development mode
```
go run main.go
```

for auto reload

```
gin -i --notifications run main.go
```

### create swagger 
```
swag init
```
 http://localhost:8080/swagger/index.html
