# Menu / Orders Service

A Go project containing an orders service that exposes a gRPC API. The `orders` application starts the gRPC server, while `common/api` contains the API definition and generated gRPC code.

## Project layout

- `orders/main.go` — Creates the gRPC server, listens on the configured address, initializes the store and service, registers the gRPC handler, and serves requests.
- `orders/` — Contains the order application code, including the store, service, and gRPC handler.
- `common/api/oms.proto` — Protocol Buffers definition for the API.
- `common/api/oms_grpc.pb.go` — Generated gRPC client/server interfaces and service registration code. Do not edit manually; regenerate it from the `.proto` file.
- `common/` — Shared project code, including `EnvString`, which reads an environment variable with a default value.

## gRPC API

The generated API defines the `OrderService` service with one RPC:

- `CreateOrder` — Accepts a `CreateOrderRequest` and returns an `Order`.

The server registers this service through `RegisterOrderServiceServer`. The generated server interface requires implementations to embed `UnimplementedOrderServiceServer` for forward compatibility.

## Configuration

The server address is read from the `GRPC_ADDR` environment variable. In the shown `orders/main.go`, its default is `localhost:20000`.

Set a different address if that port is unavailable.

**Git Bash:**

```bash
GRPC_ADDR=localhost:50051 go run .
```

**PowerShell:**

```powershell
$env:GRPC_ADDR = "localhost:50051"
go run .
```

## Run locally

From the project root, download dependencies:

```bash
go mod tidy
```

Then start the orders server:

```bash
cd orders
go run .
```

If Air is installed and configured for the `orders` directory, you can use:

```bash
air
```

The server logs its listening address when it starts. A “failed to listen” error means it could not bind to the configured address; try another port or check whether the port is already in use.

## Dependencies

The project uses Go modules and gRPC-Go (`google.golang.org/grpc`). Keep the module dependencies up to date with `go mod tidy`.

## Startup behavior

The current `orders/main.go` also calls `svc.CreateOrder(context.Background())` during startup. This means the service’s `CreateOrder` method is called when the server launches, in addition to registering the gRPC handler.