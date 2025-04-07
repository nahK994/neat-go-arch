## 🧼 Clean Architecture Overview

Clean Architecture is all about organizing our project in a way that’s:
- 🛠️ **Easy to maintain**
- 🧪 **Easy to test**
- 🔌 **Easy to integrate new technology**

It achieves this by:  
**Separating concerns** — Dividing the application into four key building blocks: Entity, Use Case, Interface, and Infrastructure.  
**Dependency Inversion Principle** — Higher-level components don’t depend on lower-level ones — they only depend on abstractions.

### ⚡ Why Clean Architecture?
- 🧠 **Decoupled:** Change the database or framework without breaking business logic.
- 🧪 **Testable:** Easily unit test use cases with mocks.
- 🌱 **Scalable:** Add new features without tangled code.
- 🧩 **Flexible:** Swap external tools easily — like switching from SQL to Redis.

<br>
<br>
Let’s break down these four building blocks, ordered from highest priority (most important) to lowest priority (least important):

#### 1️⃣ Entity (Core Business Models) 🧠
- 💖 **The heart of the application**, containing core business objects and basic validation rules.
- Should not depend on any external tech (like databases, web frameworks, or APIs).
- **Example:** `User` model enforcing business rules like:
   - ✅ Users under 18 cannot create an account
   - 🚫 User email must be unique
- **Important Note:** Field-level checks like email format or age validation belong here. But checking email uniqueness belongs in the use case layer, as it requires interacting with external systems (like a database).

#### 2️⃣ Use Cases (Application Logic) ⚙️🧠
- **Define the app’s behavior** — they’re the brains of the operation.
- Orchestrate how entities interact and call interfaces for external actions.
- Have no knowledge of HTTP, databases 🗄️, or external systems — only pure business logic.
- **Example:** `CreateUser`, `DeleteUser`, `Login`.

#### 3️⃣ Interface (External Communication) 🌐📡
- **Connect the outside world** to the app.
- Translate external requests into something the use cases understand and send proper responses back.
- Includes:
  - 🧑‍✈️ Handlers (Gin)
  - 🧼 Middleware (JWT based authentication, authorization)
- **Example:**  
   `user_handler.go` manages routes like `POST /signup`, `POST /login`.

#### 4️⃣ Infrastructure (External Stuff) 🏗️🔧
- **External dependencies** — databases, message queues, caching systems, third-party services.
- It’s a low-level detail — the use cases should never depend directly on infrastructure.
- **Example:**  
   `user_repo.go` handles SQL queries. `migrate.go` sets up DB.

---

Let’s see Clean Architecture in action with a simple CRUD project.

### 🗂️ `simple-CRUD` Folder Structure
```
📂 simple-CRUD
├── 📂 cmd
│   └── main.go                  # Entry point
│
├── 📄 docker-compose.yml        # Docker setup
├── 📄 Makefile                  # Task runner
├── 📄 go.mod / go.sum           # Dependencies
│
└── 📂 pkg
    ├── 📂 app                   # App config
    │   └── config.go
    │
    ├── 📂 entity                # Core business types
    │   └── type.go
    │
    ├── 📂 errors                # Custom error types
    │   └── errors.go
    │
    ├── 📂 handler               # HTTP handlers (Gin)
    │   └── user_handler.go
    │
    ├── 📂 middleware            # JWT auth, logging, etc.
    │   └── middleware.go
    │
    ├── 📂 repository            # Database interactions
    │   ├── migrate.go
    │   ├── migrations
    │   │   └── 00001_create_user_table.sql
    │   └── user_repo.go
    │
    ├── 📂 router                # Route definitions
    │   └── route.go
    │
    └── 📂 usecase              # Business logic
        ├── helper.go           # Token generation, hashing
        └── user_usecase.go     # Signup, login, user logic          
```

---

### 🧰 Prerequisites

Make sure you've got these tools installed first:

- 🐳 **Docker** — for spinning up Postgres
- ⚙️ **Make** — to run tasks with a single command  
- 🍸 **Go + Gin** — the fast HTTP web framework

---

### 🧠 How the Layers Interact
The typical flow follows:
```
[HTTP Request] 
    ↓
[Middleware]        // 🔐 Authentication & Authorization
    ↓
[Handler]           // 🎯 Route handling (only gets here *if* auth passes)
    ↓
[Use Case]          // 🧠 Executes business logic (assumes user is already validated)
    ↓
[Repository]        // 🗄️ Interacts with DB
    ↓
[Database]          // 🧬 Actual storage

```
💡 **Why does `usecase -> repository` happen?**

As we discussed before, higher-level components don’t depend on lower-level ones. That’s why **use cases depend on an abstraction (interface)**, not the concrete implementation of the repository. This keeps the dependency direction correct:
- The **use case defines an interface** like `UserRepository`, describing required operations.
- The **repository layer implements this interface**, handling actual database operations.
- The **use case knows only the interface**, not the implementation — maintaining decoupling.

---

### 📝 Final Thoughts

Clean Architecture might seem like overkill for small projects, but it pays off big time for larger, long-term projects. By keeping business logic separate from tech details, you avoid tangled dependencies and make your app more testable, flexible, and maintainable.

When in doubt, keep your code clean and your architecture cleaner! 🧼✨
