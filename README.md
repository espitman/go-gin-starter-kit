# go-gin-starter-kit

The fast way to create a restful apis with Gin Framework with a structured project that defaults to mongodb and redis and rabbitMQ and cron and elastic search

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

## elastic search

### sample

```
type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

elk.CreateIndex("tweets")
tweet := Tweet{User: "username", Message: "message"}
result := elk.AddData("tweets", "doc", tweet)
```

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

### create consumer

```
go run main.go generate consumer {$name}
```

### create cron

```
go run main.go generate cron {$name}
```
