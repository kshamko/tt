.PHONY: gen-cover cover cover-func
PACKAGES=`go list ./... | grep -v vendor`
SWAGGER := docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(shell pwd)/internal quay.io/goswagger/swagger
LINTER := docker run --rm -v $(shell pwd):/app -v ${GOPATH}/pkg/mod:/go/pkg/mod -w /app golangci/golangci-lint golangci-lint run --enable-all


swagger:
	docker pull quay.io/goswagger/swagger
	cd internal && $(SWAGGER) generate server --spec ../api/api.yaml --exclude-main && cd ..

lint:
	docker pull golangci/golangci-lint	
	$(LINTER)

test:
	go test ./...

server:
	docker pull namely/protoc
	docker run --rm -v $(shell pwd):/defs namely/protoc -I. --openapiv2_out=. --grpc-gateway_out=logtostderr=true:. --go_out=plugins=grpc:. api/grpc.proto