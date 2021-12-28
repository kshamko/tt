FROM golang:1.17-alpine as build

WORKDIR /app
ENV GO111MODULE=on

RUN apk --no-cache add make git ca-certificates gcc g++

COPY . .
RUN go mod download
RUN make test
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bin/app cmd/main.go

FROM scratch
COPY --from=build /bin/app /
ENTRYPOINT ["/app"]
