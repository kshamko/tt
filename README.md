## Build & RUN

*IMPORTANT!!!* before running the services please add json file with data to gateway/assets/data folder

```bash
$ git clone git@github.com:kshamko/tt.git
$ docker-compose up
```

The application will be started on port 8080 and it will be possible to request it localy:
```bash
curl -XGET 'http://localhost:8080/api/v1/data/{id}'
```

or in a browser

## TODO list

- add trace_id to logs
- think on how to improve handling of coordinates as floats are skewed during encoding to transfer via grpc
- add validation of input parameters (at least for coordinates)
- fix linter errors and linter check to the Dockerfiles
- improve errors output to users by the gateway
- use grpc stream to send data to internal service
- would move data import from gateway to internal service