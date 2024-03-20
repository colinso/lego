# lego
Lego is a tool for limiting the writing of boilerplate code and copypasta when creating an application. It provides functionality to generate common necessary structures such as HTTP handlers, routers, models, configurations, and more.

## Getting Started

To install, run:
```
go install github.com/colinso/lego
```

Create a config file. Currently only yaml is supported. Example can be found in `test.yaml`.

To generate a basic Go microservice, run `lego generate <pathToConfig> <pathToProject>`. This will create your router, server, handlers, and set up dependency injection and config files.

Run your service with `export APP_ENV=dev && go run main.go` from the project root. This will start an HTTP server at port 8080 by default.

## Roadmap
- Bugs
    - Can't guarantee order of maps, which breaks logic layer
- Generate:
    - Logic/service layer
- Use swaggo to generate swagger docs
- Allow for default service configs
- Outline http responses and corresponding errors
    - Outline errors to create
- Ensure we're handling concurrent requests
    - https://medium.com/insiderengineering/concurrent-http-requests-in-golang-best-practices-and-techniques-f667e5a19dea
    - Implement a worker pool
- AWS Lambda support
- Dockerization
- Create project in provided directory by default. Create new directory with a certain flag
- Add a makefile for ease
- cli tool support
- Only overwrite existing files if specified --> Ask for confirmation?
- Error handling
