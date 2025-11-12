# 🚀 gRPC Sample - Protocol Buffers Repository

**Shared Protocol Buffer definitions dan Generated Code untuk gRPC Microservices Ecosystem**

[![Go](https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Protocol Buffers](https://img.shields.io/badge/Protocol_Buffers-v3-4285F4?style=flat&logo=protobuf)](https://protobuf.dev/)
[![Buf](https://img.shields.io/badge/Buf-v2-2496ED?style=flat&logo=buf)](https://buf.build/)
[![gRPC Gateway](https://img.shields.io/badge/gRPC_Gateway-v2.27.3-green?style=flat)](https://github.com/grpc-ecosystem/grpc-gateway)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-v2-85EA2D?style=flat&logo=swagger)](https://swagger.io/)

## 📋 Deskripsi

Repository ini adalah **central source of truth** untuk semua Protocol Buffer definitions dan generated code yang digunakan dalam gRPC microservices ecosystem. Repository ini menyediakan **shared contracts** antara services, dengan automatic code generation untuk Go dan OpenAPI/Swagger documentation.

Cocok sebagai **referensi portofolio** untuk mendemonstrasikan kemampuan dalam:
- 🎯 **Protocol Buffers Design** - Modern proto3 syntax dengan validation
- 🔄 **Code Generation** - Multi-language support dengan Buf
- 🌐 **gRPC Gateway** - HTTP REST API mapping dari gRPC services  
- 📚 **API Documentation** - Automatic OpenAPI/Swagger generation
- 🛡️ **Schema Evolution** - Breaking change detection dan versioning

## ✨ Fitur Berdasarkan Struktur Project

### 🔧 **Protocol Buffer Services**

**1. Hello Service (hello/v1/)**
- ✅ **4 gRPC Method Types** - Unary, Server Stream, Client Stream, Bidirectional Stream
- ✅ **Input Validation** - buf.validate dengan min/max length constraints
- ✅ **HTTP REST Mapping** - gRPC Gateway annotations untuk REST 

**2. Config Service (config/v1/)**
- ✅ **OpenAPI Configuration** - Global swagger definitions
- ✅ **Security Definitions** - API key authentication schema
- ✅ **API Metadata** - Title, version, dan documentation setup

### 🛠️ **Code Generation Pipeline**

**Buf Configuration (buf.yaml v2):**
- 📦 **Dependencies Management** - googleapis, protovalidate, grpc-gateway
- 🔍 **Linting** - STANDARD rules untuk code quality
- 🛡️ **Breaking Change Detection** - FILE-level breaking change protection

**Generation Targets (buf.gen.yaml):**
```yaml
✅ Go Protocol Buffers (v1.36.10)
✅ Go gRPC Stubs (v1.5.1)  
✅ gRPC Gateway (v2.27.3) - HTTP REST mapping
✅ OpenAPI/Swagger (v2.27.3) - API documentation
🚧 PHP Support (commented - future expansion)
```

## 🏗️ Struktur Project (Actual)

```
grpc-sample/
├── proto/                          # Source Protocol Buffer definitions
│   ├── config/v1/
│   │   └── config.proto           # OpenAPI configuration & security schema
│   ├── hello/v1/
│   │   ├── hello.proto            # Message definitions dengan validation
│   │   └── service.proto          # 4 gRPC methods + HTTP mappings
│   └── resiliency/v1/
│       ├── resiliency.proto       # Fault injection messages
│       └── service.proto          # Resiliency testing methods
├── protogen/                       # Generated Go code (auto-generated)
│   ├── config/v1/
│   ├── hello/v1/
│   │   ├── hello.pb.go           # Protocol Buffer messages
│   │   ├── service.pb.go         # Service definitions  
│   │   ├── service_grpc.pb.go    # gRPC server/client stubs
│   │   └── service.pb.gw.go      # gRPC Gateway HTTP handlers
│   └── resiliency/v1/
├── openapiv2/                      # Generated OpenAPI documentation
│   ├── embed.go                   # Go embed untuk swagger files
│   └── openapiv2.swagger.json     # Merged OpenAPI specification
├── buf.yaml                        # Buf configuration v2
├── buf.gen.yaml                   # Code generation configuration
├── buf.lock                       # Dependency lock file
├── Makefile                       # Build automation
└── go.mod                         # Go module (v1.24.4)
```

## 🚀 Quick Start

### Prerequisites
- Go 1.24.4+
- Buf CLI v1.59.0+

### 🔧 Setup & Generation

```bash
# Clone repository
git clone https://github.com/achtarudin/grpc-sample.git
cd grpc-sample

# Install Buf CLI
make install-tools

# Install Go dependencies
make install-deps

# Generate code dari proto files
make generate-proto

# Clear generated files (jika perlu reset)
make clear-protogen
```

### 📦 Penggunaan sebagai Dependency

**Go Module Import:**
```go
// Import generated Go code
import (
    hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
    configv1 "github.com/achtarudin/grpc-sample/protogen/config/v1"
    resiliencyv1 "github.com/achtarudin/grpc-sample/protogen/resiliency/v1"
)
```

**go.mod dependency:**
```go
require github.com/achtarudin/grpc-sample v0.1.3
```

## 📚 API Documentation

### gRPC Methods dengan HTTP Mapping

**1. SayHello (Unary)**
```protobuf
// gRPC: HelloService/SayHello
// HTTP: POST /hello/v1/say-hello
rpc SayHello(SayHelloRequest) returns (SayHelloResponse)
```

**2. SayManyHellos (Server Streaming)**
```protobuf  
// gRPC: HelloService/SayManyHellos  
// HTTP: POST /hello/v1/say-many-hellos
rpc SayManyHellos(SayManyHellosRequest) returns (stream SayManyHellosResponse)
```

**3. SayHelloToEveryone (Client Streaming)**
```protobuf
// gRPC: HelloService/SayHelloToEveryone
// HTTP: POST /hello/v1/say-hello-to-everyone  
rpc SayHelloToEveryone(stream SayHelloToEveryoneRequest) returns (SayHelloToEveryoneResponse)
```

**4. SayHelloContinuous (Bidirectional Streaming)**
```protobuf
// gRPC: HelloService/SayHelloContinuous
// HTTP: POST /hello/v1/say-hello-continuous
rpc SayHelloContinuous(stream SayHelloContinuousRequest) returns (stream SayHelloContinuousResponse)
```

## 🛠️ Development Features

### Available Make Commands:
```bash
make install-tools     # Install Buf CLI v1.59.0
make install-deps      # Update Go dependencies
make generate-proto    # Generate semua code dari proto files
make clear-protogen    # Clean generated files
```

### Buf Features:
- **Dependency Management** - Auto-update external dependencies
- **Linting** - STANDARD proto linting rules
- **Breaking Changes** - FILE-level change detection
- **Multi-target Generation** - Go, gRPC Gateway, OpenAPI dalam satu command

## 🔄 Workflow Integration

Repository ini digunakan sebagai **shared dependency** dalam:

1. **gRPC Server** - Import generated Go stubs untuk server implementation
2. **gRPC Client** - Import generated Go stubs untuk client calls  
3. **gRPC Gateway** - Import gateway handlers untuk HTTP REST API
4. **API Documentation** - Serve embedded OpenAPI specification

## 🤝 Kontribusi

Repository ini adalah **core component** untuk gRPC microservices ecosystem. Contributions welcome untuk:
- 🔄 **New Services** - Tambah proto definitions baru
- 📚 **Documentation** - Improve proto comments dan annotations
- 🛡️ **Validation Rules** - Enhanced buf.validate constraints
- 🌐 **Multi-programming Support** - Generate code untuk bahasa lain

## 👨‍💻 Author

**Achmad Ach Tarudin**
- 🌐 Implementation: [grpc-server.cutbray.tech](https://grpc-server.cutbray.tech)
- 🐙 GitHub: [@achtarudin](https://github.com/achtarudin)
- 📦 Module: `github.com/achtarudin/grpc-sample`

---

> **💼 Catatan Portofolio:** Repository ini mendemonstrasikan expertise dalam **Protocol Buffers design**, **code generation automation**, **API contract management**, dan **microservices architecture**. Menunjukkan kemampuan gRPC ecosystems dengan **proper versioning**, **breaking change protection**.
