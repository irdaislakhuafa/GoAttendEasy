# GoAttendEasy

GoAttendEasy is an application to make attendance management easier with a series of simple features to manage attendance more efficiently

## Database Schema

Below is database schema for this app

- [schema](https://dbdiagram.io/d/GoAttendEasy-62500b892514c97903f5e23d)

# The Technology Stack Or Library Used In Developing This application

- [gocron](https://github.com/go-co-op/gocron.git): A Golang Job Scheduling Package
- [ent](https://github.com/ent/ent.git): An Entity Framework For Go
- [echo](https://github.com/labstack/echo.git): High performance, minimalist Go web framework

# Published App

This application is a Rest API type with a URL prefix which has been published on fly.io below

- https://go-attend-easy.fly.dev

You can use Rest API usage with postman documentation below and select environment `PreTest - PROD`

- https://www.postman.com/orange-eclipse-673431/workspace/digitels/documentation/17180185-9f803539-73da-4860-800b-82f722cc3f2b

# Run App On Local

If you want to run this app on local, you can use docker and docker-compose to make it more easy without any configuration needed

```bash
$ docker-compose up -d
```

or if you have `Go` installed, you can run command belo

```go
$ go run src/cmd/main.go -env prod
```

user `prod` env to load config for production in local and user `local` env to load config for local app
