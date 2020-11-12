# go-gin-starter-kit

The fast way to create a restful apis with Gin Framework with a structured project that defaults to mongodb and redis and rabbitMQ

## how to start

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

### create model

```
go run main.go generate model {$name}
```

### create dto

```
go run main.go generate dto {$name}
```

## RabbitMQ

### publish message

first create publisher (on init of package)

```
rabbitmq.CreatePublisher(exchangeName string, exchangeType string, durable bool, queueName string)
```

next publish message

```
rabbitmq.Publish(exchangeName string, queueName string, body string)
```
