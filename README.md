## 📋 项目描述

**医疗预约微服务系统** - 基于 Go 语言开发的分布式医疗服务平台，提供医生排班、患者预约、科室管理等核心功能，采用 gRPC 进行服务间通信，包含独立的 API 网关层。

## 🏗️ 系统架构
- **API Gateway** (`api-gateway/`)：统一入口，处理客户端请求的路由和转发
- **Medi-Serve** (`medi-serve/`)：核心医疗服务，处理业务逻辑
- **Common** (`common/`)：共享配置、数据模型和工具包
- **Proto** (`proto/`)：gRPC 协议定义文件

---

# 🏥 医疗预约微服务系统

一个基于 Go 语言和 gRPC 的分布式医疗预约平台，支持医生排班、患者预约、科室管理等功能。

## 🏗️ 项目结构

```
.
├── .gitignore
├── README.md
├── api-gateway/                    # API 网关服务
│   ├── basic/cmd/main.go          # 网关启动入口
│   ├── handler/                   # 请求处理
│   │   ├── api/medi.go           # API 处理器
│   │   └── request/medi.go       # 请求封装
│   └── router/router.go           # 路由配置
├── common/                        # 公共模块
│   ├── config/                   # 配置文件管理
│   │   ├── config.go
│   │   ├── config.yaml
│   │   └── global.go
│   ├── init/init.go              # 初始化逻辑
│   ├── models/                   # 数据模型
│   │   ├── appointment.go       # 预约模型
│   │   ├── department.go        # 科室模型
│   │   ├── doctor.go            # 医生模型
│   │   ├── patient.go           # 患者模型
│   │   └── schedule.go          # 排班模型
│   └── pkg/                      # 工具包
│       ├── alipay.go            # 支付宝支付
│       ├── genware.go           # 通用中间件
│       └── page.go              # 分页工具
├── medi-serve/                   # 医疗服务
│   ├── basic/cmd/main.go        # 服务启动入口
│   └── handler/service/service.go # 业务服务实现
└── proto/                        # gRPC 协议定义
    └── medi/
        ├── medi.proto           # 协议定义
        ├── medi.pb.go           # 生成的 Go 代码
        └── medi_grpc.pb.go      # 生成的 gRPC 代码
```

## ✨ 功能特性

### 核心业务模块
- **患者管理**：患者注册、信息维护
- **医生管理**：医生信息、资质管理
- **科室管理**：医院科室分类与配置
- **排班管理**：医生出诊时间安排
- **预约系统**：在线预约、取消、改期

### 技术特性
- **微服务架构**：API 网关 + 业务服务分离
- **gRPC 通信**：高性能服务间通信
- **统一配置管理**：YAML 配置文件中心化管理
- **支付宝支付集成**：支持在线支付功能
- **模型驱动开发**：清晰的领域模型定义

## 🛠️ 技术栈

- **语言**: Go 1.19+
- **通信协议**: gRPC + Protocol Buffers
- **API 网关**: 自定义网关层
- **配置管理**: YAML
- **支付集成**: 支付宝 SDK
- **架构模式**: 微服务架构

## 🚀 快速开始

### 1. 克隆项目
```bash
git clone <repository-url>
cd medical-appointment-system
```

### 2. 安装依赖
```bash
# 安装 Go 依赖
go mod download

# 安装 Protocol Buffers 编译器 (如果需要重新生成)
# 参考: https://grpc.io/docs/protoc-installation/
```

### 3. 生成 gRPC 代码 (可选)
```bash
cd proto/medi
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       medi.proto
```

### 4. 配置系统
编辑配置文件：
```bash
cp common/config/config.yaml.example common/config/config.yaml
# 修改数据库、服务端口、支付宝等配置
```

### 5. 启动服务

#### 启动医疗服务：
```bash
cd medi-serve/basic/cmd
go run main.go
```

#### 启动 API 网关：
```bash
cd api-gateway/basic/cmd
go run main.go
```

### 6. 访问服务
- API 网关默认端口：`8080`
- gRPC 服务端口：`50051` (根据配置调整)

## 📡 API 接口

### 网关路由示例
```
GET    /api/doctors          # 获取医生列表
POST   /api/appointments     # 创建预约
GET    /api/departments      # 获取科室列表
PUT    /api/schedules/{id}   # 更新排班
```

### gRPC 服务
- 服务名称：`MediService`
- 定义文件：`proto/medi/medi.proto`

## 🔧 开发指南

### 添加新的 API
1. 在 `proto/medi/medi.proto` 中添加新的 RPC 方法
2. 重新生成 gRPC 代码
3. 在 `medi-serve/handler/service/service.go` 中实现服务
4. 在 `api-gateway/handler/api/medi.go` 中添加 HTTP 接口处理
5. 在 `api-gateway/router/router.go` 中注册路由

### 数据模型扩展
在 `common/models/` 目录下添加新的 `.go` 文件定义模型结构。

## 📊 数据库

系统使用关系型数据库（如 MySQL/PostgreSQL）存储业务数据。具体数据库配置在 `common/config/config.yaml` 中设置。

## 🧪 测试

运行测试：
```bash
go test ./...
```

## 📄 许可证

MIT License

---

这个 README.md 清晰地展示了项目的微服务架构、功能模块和启动方式，适合开发者和部署人员使用。