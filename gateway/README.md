# Description

HTTP gateway

## Approach

It is important to have API documented, thus there is a problem of documentation and API implementation consistency. To solve the proble the following approach was used:

1. First step is to create swagger spec to describe the API (**api/gw.swagger.yaml**). The naming of the entities is kind of abstract (i.e Data)  to prevent exposing this test task to others via GitHub search.
2. Use go-swagger to generate server's code:
```bash
$ make swagger
```
3. Implement endpoints' handlers
4. If there is a requirement to change something in API, start from 1.

## Service Endpoints

Also an additionl port is exposed for debug/healthcheck/swagger-ui purposes

1. Metrics: http://localhost:2112/metrics
2. Healthcheck: http://localhost:2112/healthz
3. Swagger-UI: http://localhost:2112/swagger-ui
4. See **internal/debug/debug.go** for more details
