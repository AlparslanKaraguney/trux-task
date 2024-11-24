# **Smart Service**

A gRPC-based service for managing smart models and features, built with Go, GORM, and PostgreSQL, and deployable on Minikube.

---

## **Features**

- **gRPC API**: Provides endpoints for managing smart models and features.
- **Unit Tested**: Comprehensive test coverage with mock generation for interfaces.
- **Minikube Deployment**: Easily deployable to a local Kubernetes cluster.
- **Structured Logging**: Logs all activities using `logrus` for better observability.
- **Database Integration**: Supports PostgreSQL with migration and schema management.

---

## **Table of Contents**

- [Setup and Development](#setup-and-development)
  - [Update Proto Files](#update-proto-files)
  - [Generate Mocks](#generate-mocks)
  - [Run Unit Tests](#run-unit-tests)
  - [Code Coverage](#code-coverage)

---

## **Setup and Development**

### **Update Proto Files**

To update the gRPC proto files, use:

```bash
protoc --go_out=./ --go-grpc_out=./ proto/smart_service.proto
```

### **Generate Mocks**

Generate mock files for interfaces using mockgen:

```bash
go generate ./...
```

### **Run Unit Tests**

Run all unit tests:

```bash
go test ./...
```

### **Code Coverage**

To generate and view the code coverage report: (These files excluded from the coverage since they are not part of the actual code or auto generated ("mocks", "client", "proto"))

```bash
go test -coverpkg=$(go list ./... | grep -v '/mocks' | grep -v /proto | grep -v /client) -coverprofile=coverage.out ./...
```

```
go tool cover -html=coverage.out -o coverage.html
```

```
start coverage.html
```

Or use the provided script. Also gives the total coverage percentage:

```
sh coverage.sh
```

```
start coverage.html
```
