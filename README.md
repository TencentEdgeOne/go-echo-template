# Go Cloud Functions on EdgeOne Pages - Echo Framework

A full-stack demo application built with Next.js + Tailwind CSS frontend and Go Echo backend, showcasing how to deploy Go Cloud Functions using the Echo framework on EdgeOne Pages with RESTful API routing.

## 🚀 Features

- **Echo Framework Integration**: High-performance, minimalist Go web framework with custom middleware and route groups
- **Modern UI Design**: Dark theme with #1c66e5 accent color, responsive layout with interactive elements
- **Interactive API Testing**: Built-in API endpoint panel — click "Call" to test each REST endpoint in real-time
- **RESTful API Design**: Complete Todo CRUD operations with structured route groups (`/api/todos`)
- **TypeScript Support**: Complete type definitions and type safety on the frontend

## 🛠️ Tech Stack

### Frontend
- **Next.js 15** - React full-stack framework (with Turbopack)
- **React 19** - User interface library
- **TypeScript** - Type-safe JavaScript
- **Tailwind CSS 4** - Utility-first CSS framework

### UI Components
- **shadcn/ui** - High-quality React components
- **Lucide React** - Beautiful icon library
- **class-variance-authority** - Component style variant management
- **clsx & tailwind-merge** - CSS class name merging utilities

### Backend
- **Go 1.21** - Cloud Functions runtime
- **Echo v4** - High-performance, minimalist Go web framework

## 📁 Project Structure

```
go-echo/
├── cloud-functions/                # Go Cloud Functions source
│   ├── api.go                     # Echo app with all REST API routes
│   ├── go.mod                     # Go module definition
│   └── go.sum                     # Go dependency checksums
├── src/
│   ├── app/                       # Next.js App Router
│   │   ├── globals.css           # Global styles (dark theme)
│   │   ├── layout.tsx            # Root layout
│   │   └── page.tsx              # Main page (API testing UI)
│   ├── components/               # React components
│   │   └── ui/                   # UI base components
│   │       ├── button.tsx        # Button component
│   │       └── card.tsx          # Card component
│   └── lib/                      # Utility functions
│       └── utils.ts              # Common utilities (cn helper)
├── public/                        # Static assets
│   ├── eo-logo-blue.svg          # EdgeOne logo (blue)
│   └── eo-logo-white.svg         # EdgeOne logo (white)
├── package.json                   # Project configuration
└── README.md                     # Project documentation
```

## 🚀 Quick Start

### Requirements

- Node.js 18+
- pnpm (recommended) or npm
- Go 1.21+ (for local development)

### Install Dependencies

```bash
pnpm install
# or
npm install
```

### Development Mode

```bash
edgeone pages dev
```

Visit [http://localhost:8088](http://localhost:8088) to view the application.

### Build Production Version

```bash
edgeone pages build
```

## 🎯 Core Features

### 1. Echo REST API Routes

All API endpoints are defined in a single `cloud-functions/api.go` file using Echo's route groups:

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/` | Welcome message with route list |
| GET | `/health` | Health check (includes Go runtime info) |
| GET | `/api/todos` | List all todos |
| POST | `/api/todos` | Create a new todo |
| GET | `/api/todos/:id` | Get todo by ID |
| PATCH | `/api/todos/:id/toggle` | Toggle todo completion |
| DELETE | `/api/todos/:id` | Delete a todo |

### 2. Interactive API Testing Panel

- 7 pre-configured API endpoint cards with "Call" buttons
- Real-time JSON response display with syntax highlighting
- POST request support with pre-filled JSON body
- Loading states and error handling

### 3. Echo Framework Convention

The Go backend uses Echo's standard patterns — route groups, context binding, and custom middleware:

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

### 4. Data Model

```go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"createdAt"`
}
```

## 🔧 Configuration

### Tailwind CSS Configuration
The project uses Tailwind CSS 4 with custom color variables:

```css
:root {
  --primary: #1c66e5;        /* Primary color */
  --background: #000000;     /* Background color */
  --foreground: #ffffff;     /* Foreground color */
}
```

### Component Styling
Uses `class-variance-authority` to manage component style variants with multiple preset styles.

## 📚 Documentation

- **EdgeOne Pages Official Docs**: [https://pages.edgeone.ai/document/go](https://pages.edgeone.ai/document/go)
- **Echo Framework**: [https://echo.labstack.com/docs](https://echo.labstack.com/docs)
- **Next.js Documentation**: [https://nextjs.org/docs](https://nextjs.org/docs)
- **Tailwind CSS Documentation**: [https://tailwindcss.com/docs](https://tailwindcss.com/docs)

## 🚀 Deployment Guide

### EdgeOne Pages Deployment

1. Push code to GitHub repository
2. Create new project in EdgeOne Pages console
3. Select GitHub repository as source
4. Configure build command: `edgeone pages build`
5. Deploy project

### Go Echo Cloud Function

Create `cloud-functions/api.go` in project root with your Echo application:

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

## Deploy

[![Deploy with EdgeOne Pages](https://cdnstatic.tencentcs.com/edgeone/pages/deploy.svg)](https://edgeone.ai/pages/new?from=github&template=go-echo)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/github/choosealicense.com/blob/gh-pages/_licenses/mit.txt) file for details.
