# Go Cloud Functions on EdgeOne Pages - Echo 框架

一个基于 Next.js + Tailwind CSS 前端和 Go Echo 后端的全栈演示应用，展示如何在 EdgeOne Pages 上使用 Echo 框架部署 Go 云函数，支持 RESTful API 路由。

## 🚀 特性

- **Echo 框架集成**：高性能、简约的 Go Web 框架，支持自定义中间件和路由组
- **现代化 UI 设计**：深色主题搭配 #1c66e5 点缀色，响应式布局配合交互元素
- **交互式 API 测试**：内置 API 端点测试面板 — 点击 "Call" 实时测试每个 REST 端点
- **RESTful API 设计**：完整的 Todo CRUD 操作，使用结构化路由组（`/api/todos`）
- **TypeScript 支持**：前端完整的类型定义和类型安全

## 🛠️ 技术栈

### 前端
- **Next.js 15** - React 全栈框架（支持 Turbopack）
- **React 19** - 用户界面库
- **TypeScript** - 类型安全的 JavaScript
- **Tailwind CSS 4** - 实用优先的 CSS 框架

### UI 组件
- **shadcn/ui** - 高质量 React 组件
- **Lucide React** - 精美的图标库
- **class-variance-authority** - 组件样式变体管理
- **clsx & tailwind-merge** - CSS 类名合并工具

### 后端
- **Go 1.21** - 云函数运行时
- **Echo v4** - 高性能、简约的 Go Web 框架

## 📁 项目结构

```
go-echo/
├── cloud-functions/                # Go 云函数源码
│   ├── api.go                     # Echo 应用，包含所有 REST API 路由
│   ├── go.mod                     # Go 模块定义
│   └── go.sum                     # Go 依赖校验和
├── src/
│   ├── app/                       # Next.js App Router
│   │   ├── globals.css           # 全局样式（深色主题）
│   │   ├── layout.tsx            # 根布局
│   │   └── page.tsx              # 主页面（API 测试界面）
│   ├── components/               # React 组件
│   │   └── ui/                   # UI 基础组件
│   │       ├── button.tsx        # 按钮组件
│   │       └── card.tsx          # 卡片组件
│   └── lib/                      # 工具函数
│       └── utils.ts              # 通用工具（cn 辅助函数）
├── public/                        # 静态资源
│   ├── eo-logo-blue.svg          # EdgeOne 标志（蓝色）
│   └── eo-logo-white.svg         # EdgeOne 标志（白色）
├── package.json                   # 项目配置
└── README.md                     # 项目文档
```

## 🚀 快速开始

### 环境要求

- Node.js 18+
- pnpm（推荐）或 npm
- Go 1.21+（本地开发需要）

### 安装依赖

```bash
pnpm install
# 或
npm install
```

### 开发模式

```bash
edgeone pages dev
```

访问 [http://localhost:8088](http://localhost:8088) 查看应用。

### 构建生产版本

```bash
edgeone pages build
```

## 🎯 核心功能

### 1. Echo REST API 路由

所有 API 端点定义在单个 `cloud-functions/api.go` 文件中，使用 Echo 的路由组：

| 方法 | 路由 | 说明 |
|------|------|------|
| GET | `/` | 欢迎消息及路由列表 |
| GET | `/health` | 健康检查（包含 Go 运行时信息） |
| GET | `/api/todos` | 获取待办事项列表 |
| POST | `/api/todos` | 创建新待办事项 |
| GET | `/api/todos/:id` | 根据 ID 获取待办事项 |
| PATCH | `/api/todos/:id/toggle` | 切换待办事项完成状态 |
| DELETE | `/api/todos/:id` | 删除待办事项 |

### 2. 交互式 API 测试面板

- 7 个预配置的 API 端点卡片，带 "Call" 按钮
- 实时 JSON 响应展示
- POST 请求支持预填充 JSON Body
- 加载状态和错误处理

### 3. Echo 框架约定

Go 后端使用 Echo 的标准模式 — 路由组、上下文绑定和自定义中间件：

```go
package main

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func main() {
    e := echo.New()

    e.GET("/health", func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "status":    "ok",
            "framework": "echo",
        })
    })

    api := e.Group("/api")
    api.GET("/todos", listTodos)
    api.POST("/todos", createTodo)
    api.GET("/todos/:id", getTodo)
    api.PATCH("/todos/:id/toggle", toggleTodo)
    api.DELETE("/todos/:id", deleteTodo)

    e.Start(":9000")
}
```

### 4. 数据模型

```go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"createdAt"`
}
```

## 🔧 配置说明

### Tailwind CSS 配置
项目使用 Tailwind CSS 4，支持自定义颜色变量：

```css
:root {
  --primary: #1c66e5;        /* 主色调 */
  --background: #000000;     /* 背景色 */
  --foreground: #ffffff;     /* 前景色 */
}
```

### 组件样式
使用 `class-variance-authority` 管理组件样式变体，支持多种预设样式。

## 📚 文档入口

- **EdgeOne Pages 官方文档**：[https://edgeone.ai/document/go-functions](https://edgeone.ai/document/go-functions)
- **Echo 框架文档**：[https://echo.labstack.com/docs](https://echo.labstack.com/docs)
- **Next.js 文档**：[https://nextjs.org/docs](https://nextjs.org/docs)
- **Tailwind CSS 文档**：[https://tailwindcss.com/docs](https://tailwindcss.com/docs)

## 🚀 部署指南

### EdgeOne Pages 部署

1. 将代码推送到 GitHub 仓库
2. 在 EdgeOne Pages 控制台创建新项目
3. 选择 GitHub 仓库作为源
4. 配置构建命令：`edgeone pages build`
5. 部署项目

### Go Echo 云函数配置

在项目根目录创建 `cloud-functions/api.go`，编写 Echo 应用：

```go
package main

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func main() {
    e := echo.New()

    e.GET("/hello", func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "message": "Hello from Go Echo on EdgeOne Pages!",
        })
    })

    e.Start(":9000")
}
```

## 部署

[![Deploy with EdgeOne Pages](https://cdnstatic.tencentcs.com/edgeone/pages/deploy.svg)](https://console.cloud.tencent.com/edgeone/pages/new?from=github&template=go-echo)

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](https://github.com/github/choosealicense.com/blob/gh-pages/_licenses/mit.txt) 文件了解详情。
