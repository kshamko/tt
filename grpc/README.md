# gRPC Service

- The service exposes endpoint to add/get data about ports
- API spec is described in *api/grpc.proto*. The naming of the entities is kind of abstract (i.e Entity)  to prevent exposing this test task to others via GitHub search.
- To generate code for the server run  
```bash
$ make server
```
- As a datastorage just a map was used *internal/datasource/map.go*. Datasource is an interface so any other implementation might be used easily
- The application exposes metrics available in prometheuse format at *http(s)://host:port/metrics* (in currrent docker compose configuration it is http://localhost:2113/metrics)
- The application exposes health checks at *http(s)://host:port/healthz* (in currrent docker compose configuration it is http://localhost:2113/healthz)
