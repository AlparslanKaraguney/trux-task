# **Smart Service API Documentation**

The **Smart Service API** is a gRPC-based service for managing smart models and their associated features. It provides endpoints for creating, retrieving, listing, updating, and deleting smart models and features.

---

## **Getting Started**

### **Requirements**

- **gRPC Client**: Install `grpcurl`, `Postman`, or use any gRPC client library (e.g., `grpc-go`, `grpc-python`).
- **Proto Files**: Obtain the `smart_service.proto` file to generate stubs for your client.

### **Endpoint**

The service listens on the following endpoints:

- **gRPC**: `localhost:50051` (default)

---

## **Service Methods**

### **1. CreateSmartModel**

Creates a new Smart Model.

#### **Request**

```proto
message SmartModelRequest {
  SmartModel model = 1;
}
```

| Field   | Type         | Description                |
| ------- | ------------ | -------------------------- |
| `model` | `SmartModel` | The smart model to create. |

#### **Response**

```proto
message SmartModelResponse {
  SmartModel model = 1;
}
```

| Field   | Type         | Description              |
| ------- | ------------ | ------------------------ |
| `model` | `SmartModel` | The created smart model. |

#### **Example gRPCurl Command**

```bash
grpcurl -plaintext -d '{
  "model": {
    "name": "Smart Watch",
    "identifier": "sw-001",
    "type": "Device",
    "category": "Wearable"
  }
}' localhost:50051 smartservice.SmartService/CreateSmartModel
```

---

### **2. GetSmartModel**

Retrieves a specific Smart Model by its id.

#### **Request**

```proto
message SmartModelQuery {
  int32 id = 1;
}
```

| Field | Type    | Description                       |
| ----- | ------- | --------------------------------- |
| `id`  | `int32` | The unique id of the smart model. |

#### **Response**

```proto
message SmartModelResponse {
  SmartModel model = 1;
}
```

| Field   | Type         | Description                |
| ------- | ------------ | -------------------------- |
| `model` | `SmartModel` | The requested smart model. |

#### **Example gRPCurl Command**

```bash
grpcurl -plaintext -d '{
  "id": 1
}' localhost:50051 smartservice.SmartService/GetSmartModel
```

---

### **3. ListSmartModels**

Lists all Smart Models with optional filters and pagination.

#### **Request**

```proto
message SmartModelListQuery {
  int32 limit = 1;
  int32 offset = 2;
  string name = 3;
  string identifier = 4;
  string type = 5;
  string category = 6;
  string orderBy = 7;
}
```

| Field        | Type     | Description                                                                |
| ------------ | -------- | -------------------------------------------------------------------------- |
| `limit`      | `int32`  | The number of records to return per page. Default:10 (optional)            |
| `offset`     | `int32`  | The starting position for pagination. Default:0 (optional)                 |
| `name`       | `string` | Filter by name (optional).                                                 |
| `identifier` | `string` | Filter by identifier (optional).                                           |
| `type`       | `string` | Filter by type (optional).                                                 |
| `category`   | `string` | Filter by category (optional).                                             |
| `orderBy`    | `string` | Field to sort results (e.g., "name"). Default:"created_at desc" (optional) |

#### **Response**

```proto
message SmartModelListResponse {
  repeated SmartModel models = 1;
}
```

| Field    | Type                  | Description               |
| -------- | --------------------- | ------------------------- |
| `models` | `repeated SmartModel` | The list of smart models. |

#### **Example gRPCurl Command**

```bash
grpcurl -plaintext -d '{
  "limit": 10,
  "offset": 0,
  "type": "Device",
  "orderBy": "name"
}' localhost:50051 smartservice.SmartService/ListSmartModels
```

---

### **4. CreateSmartFeature**

Creates a new Smart Feature for a Smart Model.

#### **Request**

```proto
message SmartFeatureRequest {
  SmartFeature feature = 1;
}
```

| Field     | Type           | Description                  |
| --------- | -------------- | ---------------------------- |
| `feature` | `SmartFeature` | The smart feature to create. |

#### **Response**

```proto
message SmartFeatureResponse {
  SmartFeature feature = 1;
}
```

| Field     | Type           | Description                |
| --------- | -------------- | -------------------------- |
| `feature` | `SmartFeature` | The created smart feature. |

#### **Example gRPCurl Command**

```bash
grpcurl -plaintext -d '{
  "feature": {
    "name": "Take Screenshot",
    "identifier": "feature-001",
    "functionality": "Captures a screenshot",
    "model_id": 1
  }
}' localhost:50051 smartservice.SmartService/CreateSmartFeature
```

---

### **5. GetSmartFeature**

Retrieves a specific Smart Feature by its id.

#### **Request**

```proto
message SmartFeatureQuery {
  int32 id = 1;
}
```

| Field | Type    | Description                         |
| ----- | ------- | ----------------------------------- |
| `id`  | `int32` | The unique id of the smart feature. |

#### **Response**

```proto
message SmartFeatureResponse {
  SmartFeature feature = 1;
}
```

| Field     | Type           | Description                  |
| --------- | -------------- | ---------------------------- |
| `feature` | `SmartFeature` | The requested smart feature. |

#### **Example gRPCurl Command**

```bash
grpcurl -plaintext -d '{
  "id": 1
}' localhost:50051 smartservice.SmartService/GetSmartFeature
```

---

## **Error Handling**

The service uses gRPC status codes for error responses:

- **`codes.InvalidArgument`**: Returned for invalid inputs.
- **`codes.NotFound`**: Returned if the requested resource does not exist.
- **`codes.Internal`**: Returned for unexpected server errors.
- **`codes.AlreadyExists`**: Returned for dublicate value errors.

---
