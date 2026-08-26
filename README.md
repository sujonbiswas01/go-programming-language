# 🚀 Go Development Setup

> A clean, practical, step-by-step guide for setting up a Go project, running it with a `dev` command, enabling automatic reload with Air, using terminal colors, understanding format specifiers, and connecting Go with PostgreSQL.

---

## 🧭 Learning Flow

```text
Go Project
    │
    ├── 01. Initialize Go Module
    │
    ├── 02. Run Go Application
    │
    ├── 03. Create "dev" Command
    │
    ├── 04. Auto Reload with Air
    │
    ├── 05. Terminal Colors
    │
    ├── 06. Format Specifiers
    │
    └── 07. PostgreSQL + pgx
```

---

# 01 — 📦 Initialize Go Project

First, create your project:

```bash
mkdir myapp
cd myapp
```

Initialize the Go module:

```bash
go mod init myapp
```

> `myapp` can be replaced with any project/module name.

Example:

```bash
go mod init my-project
```

After running the command, Go creates:

```text
myapp/
└── go.mod
```

---

# 02 — ▶️ Run the Go Application

Create a `main.go` file:

```text
myapp/
├── go.mod
└── main.go
```

Add:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go 🚀")
}
```

Run the application:

```bash
go run .
```

### What does `go run .` mean?

```text
go run
   ↓
Run the current Go package
```

This is useful when your project contains multiple `.go` files.

---

# 03 — ⚡ Create a `dev` Command

If you want a simple development command like:

```bash
dev
```

you can define it with a `Makefile`.

Create:

```text
Makefile
```

Add:

```makefile
dev:
	go run .
```

Now run:

```bash
make dev
```

### Development Flow

```text
make dev
   ↓
go run .
   ↓
Application starts
```

> **Note:** Go's `go.mod` does not have an npm-style `"scripts"` section. A `Makefile` can be used to create commands such as `make dev`.

---

# 04 — 🔄 Auto Run After Saving Go Files

Instead of manually running:

```bash
go run .
```

every time you change your code, use **Air**.

Air watches your Go files and automatically rebuilds/restarts the application when you save.

---

## Step 1 — Install Air

Run:

```bash
go install github.com/air-verse/air@latest
```

---

## Step 2 — Go to Your Project

```bash
cd your-project
```

Your project should contain a `main.go` file:

```text
your-project/
├── go.mod
└── main.go
```

---

## Step 3 — Run Air

```bash
air
```

Now save your Go file.

Air automatically performs:

```text
Save File
    ↓
Detect Change
    ↓
Build
    ↓
Restart
    ↓
Run Application
```

### Development Experience

```text
┌───────────────┐
│  Write Code   │
└───────┬───────┘
        ↓
┌───────────────┐
│  Save File    │
└───────┬───────┘
        ↓
┌───────────────┐
│ Air Detects   │
│    Change     │
└───────┬───────┘
        ↓
┌───────────────┐
│ Build + Run   │
└───────────────┘
```

---

# 05 — 🎨 Color Package

For colorful terminal output, use:

**fatih/color**

Repository:

```text
https://github.com/fatih/color
```

Install it:

```bash
go get github.com/fatih/color
```

Use it:

```go
package main

import "github.com/fatih/color"

func main() {
    color.Green("Success! 🚀")
    color.Red("Error! ❌")
    color.Yellow("Warning! ⚠️")
}
```

### Output

```text
Success! 🚀
Error! ❌
Warning! ⚠️
```

---

# 06 — 📝 Format Specifiers in Go

Go's `fmt` package provides format specifiers for formatted output.

## Common Format Specifiers

| Specifier | Used For                | Example                     |
| --------- | ----------------------- | --------------------------- |
| `%v`      | Default value           | `fmt.Printf("%v", value)`   |
| `%+v`     | Struct with field names | `fmt.Printf("%+v", user)`   |
| `%#v`     | Go representation       | `fmt.Printf("%#v", user)`   |
| `%T`      | Value type              | `fmt.Printf("%T", value)`   |
| `%s`      | String                  | `fmt.Printf("%s", name)`    |
| `%d`      | Integer                 | `fmt.Printf("%d", age)`     |
| `%f`      | Float                   | `fmt.Printf("%f", price)`   |
| `%.2f`    | Float with 2 decimals   | `fmt.Printf("%.2f", price)` |
| `%t`      | Boolean                 | `fmt.Printf("%t", active)`  |
| `%c`      | Character               | `fmt.Printf("%c", char)`    |
| `%p`      | Pointer address         | `fmt.Printf("%p", ptr)`     |

### Example

```go
package main

import "fmt"

func main() {
    name := "Sujon"
    age := 22
    price := 99.99
    active := true

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Price: %.2f\n", price)
    fmt.Printf("Active: %t\n", active)
    fmt.Printf("Type: %T\n", age)
}
```

---

# 07 — 🐘 Go + PostgreSQL

Now we can connect our Go application with PostgreSQL.

For PostgreSQL, we will use:

**pgx**

Repository:

```text
https://github.com/jackc/pgx
```

---

## Step 1 — Install pgx

Inside your Go project:

```bash
go get github.com/jackc/pgx/v5
```

Go will add the dependency to:

```text
go.mod
go.sum
```

Project:

```text
myapp/
├── go.mod
├── go.sum
└── main.go
```

---

## Step 2 — PostgreSQL Connection

Import `pgx`:

```go
import "github.com/jackc/pgx/v5"
```

Basic connection:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/jackc/pgx/v5"
)

func main() {
    ctx := context.Background()

    conn, err := pgx.Connect(
        ctx,
        "postgres://postgres:password@localhost:5432/mydb",
    )

    if err != nil {
        log.Fatal(err)
    }

    defer conn.Close(ctx)

    fmt.Println("PostgreSQL connected successfully 🚀")
}
```

---

# 🔐 PostgreSQL Connection Flow

```text
Go Application
      │
      ↓
    pgx
      │
      ↓
PostgreSQL
      │
      ↓
   Database
```

The connection URL follows this general structure:

```text
postgres://username:password@host:port/database
```

Example:

```text
postgres://postgres:password@localhost:5432/mydb
```

---

# 🧩 Complete Setup

After completing all steps, your basic project can look like:

```text
myapp/
│
├── main.go
├── go.mod
├── go.sum
└── Makefile
```

### Main commands

```bash
# Initialize
go mod init myapp

# Run
go run .

# Development command
make dev

# Install Air
go install github.com/air-verse/air@latest

# Start auto reload
air

# Install color package
go get github.com/fatih/color

# Install PostgreSQL driver
go get github.com/jackc/pgx/v5
```

---

# 🚀 Final Workflow

```text
        CREATE PROJECT
              │
              ▼
       go mod init myapp
              │
              ▼
          main.go
              │
              ▼
          go run .
              │
              ▼
       Install Air
              │
              ▼
             air
              │
              ▼
       Auto Reload 🔄
              │
              ▼
      Add Color Package 🎨
              │
              ▼
    Learn Format Specifiers 📝
              │
              ▼
       Install pgx 📦
              │
              ▼
      PostgreSQL 🐘
```

---

## 🎯 Quick Reference

| Task               | Command                                      |
| ------------------ | -------------------------------------------- |
| Initialize project | `go mod init myapp`                          |
| Run project        | `go run .`                                   |
| Run dev command    | `make dev`                                   |
| Install Air        | `go install github.com/air-verse/air@latest` |
| Start Air          | `air`                                        |
| Install color      | `go get github.com/fatih/color`              |
| Install pgx        | `go get github.com/jackc/pgx/v5`             |

---

> **Go → Run → Auto Reload → Color → Format → PostgreSQL 🚀**

**Build it. Understand it. Ship it.**
