## Update proto files using:

```
protoc --go_out=./ --go-grpc_out=./ proto/smart_service.proto

```

# Generate Mocks

````
mockgen -source=internal/storage/storage.go -destination=internal/mocks/mock_storage.go -package=mocks
```

# Run Unit Tests

```
go test ./internal/tests/...
```

## Code Coverage

To see coverage in all packages use:

```
go test -cover ./...
go test -coverpkg=$(go list ./... | grep -v '/mocks' | grep -v /proto) -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
start coverage.html
```

Or simply:
```
sh coverage.sh
start coverage.html
```
````

## Minikube Deployment

Start minikube

```
minikube start
```

Point shell to minikube docker-daemon

```
eval $(minikube docker-env)
```

Build docker image in minikube

```
docker build -t smart-service:latest .
```

Run "runs.sh" file to apply yml files

```
sh kubernetes/run.sh
```

For testing purposes start a tunnel with minikube so we can reach services from our localhost

```
minikube tunnel
```
