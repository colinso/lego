# lego
Lego is a tool for limiting the writing of boilerplate code and copypasta when creating an application. It provides functionality to generate common necessary structures such as HTTP handlers, routers, models, configurations, and more.

## Getting Started

To install, run:
```
go get github.com/colinso/lego
go install github.com/colinso/lego
```

## Roadmap
- Bugs
    - Can't guarantee order of maps, which breaks logic layer
- Generate:
    - Logic/service layer
- Use swaggo to generate swagger doc
- Allow for default service configs
- Outline http responses and corresponding errors
    - Outline errors to create
- Ensure we're handling concurrent requests
- AWS Lambda support
- Dockerization
