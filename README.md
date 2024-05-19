# lego
Lego is a tool for limiting the writing of boilerplate code and copypasta when creating an application. It provides functionality to generate common necessary structures such as HTTP handlers, routers, models, configurations, and more.

## Getting Started

To install, run:
```
go install github.com/colinso/lego
```

Create a config file. Currently only yaml is supported. Examples can be found in `yamlConfigs`.

To generate a basic Go microservice, run `lego generate <pathToConfig> <pathToProject>`. This will create your router, server, handlers, and set up dependency injection and config files.

Run your service with `make run` from the project root. This will start an HTTP server at port 8080 by default.

## Roadmap

### Quality of Life
- [ ] Add error handling so that it's clear what bad configs caused what
- [ ] Use swaggo to generate swagger docs
- [ ] Allow for default service configs - if you don't specify, what do we spin up?
- [ ] Only overwrite existing files if specified --> Ask for confirmation?
- [ ] JSON/TOML support

### HTTP
- [ ] Outline http responses and corresponding errors
- [ ] Ensure we're handling concurrent requests
    - https://medium.com/insiderengineering/concurrent-http-requests-in-golang-best-practices-and-techniques-f667e5a19dea
    - Implement a worker pool

### Repo
- [ ] Generate repo layer for DB schema
- [ ] Call repo layer from service layer

### Service Types
- [ ] AWS Lambda support
- [ ] CLI tool support

### Questions/Ideas
- What if I changed the yaml structure? Create service defined by model rather than layers?