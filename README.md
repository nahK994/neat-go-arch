## 🧼 Clean Architecture Overview

Clean Architecture is all about organizing our project in a way that’s:
- 🛠️ **Easy to maintain**
- 🧪 **Easy to test**
- 🔌 **Easy to integrate new technology**

It achieves this by **separating concerns** and following **strict dependency rules**. Higher-level components don’t depend on lower-level ones — they only depend on **abstractions**. Let’s break down the four key building blocks, ordered from **highest priority (most important)** to **lowest priority (least important)**:

### 1️⃣ Entity (Core Business Models) 🧠
- 💖 **The heart of the application**, containing **core business objects** and **business rules**.
- Should **not depend on any external tech** (like databases, web frameworks, or APIs).
- **Example:** `User` model enforcing business rules like:
   - ✅ **User email must be unique**
   - 🚫 **Users under 18 cannot create an account**
- **Important Note:** Field-level checks like **email format** and **age validation** belong here. But checks like **ensuring the email is unique** belong in the **use case layer**, as they require interacting with external systems (like a database).

### 2️⃣ Use Cases (Application Logic) ⚙️🧠
- **Define the app’s behavior** — they’re the **brains of the operation**.
- **Orchestrate how entities interact** and call interfaces for external actions.
- Have **no knowledge of HTTP, databases, or external systems** — only pure business logic.
- **Example:** `CreateUser`, `CreditAccount`, `TransferMoney`.

### 3️⃣ Interface (External Communication) 🌐📡
- **Connect the outside world** to the app — HTTP handlers, CLI commands, gRPC, etc.
- **Translate external requests** into something the **use cases understand** and send proper responses back.
- **Example:** Gin HTTP handlers, middleware, or an API gateway.

### 4️⃣ Infrastructure (Tech & Tools) 🏗️🔧
- **External dependencies** — databases, message queues, caching systems, third-party services.
- It’s a **low-level detail** — the use cases **should never depend directly on infrastructure**.
- **Example:** PostgreSQL repo, Redis cache, AWS SES for emails.

---

### ⚡ Why Clean Architecture?
- 🧠 **Decoupled:** Change the database or framework without breaking business logic.
- 🧪 **Testable:** Easily unit test use cases with mocks.
- 🌱 **Scalable:** Add new features without tangled code.
- 🧩 **Flexible:** Swap external tools easily — like switching from SQL to Redis.

---

### 🗂️ Folder Structure
```
📂 simple-CRUD
├── 📂 cmd                    # App entry point
│   └── main.go               # Starts the Gin server
│
└── 📂 pkg                    # Reusable app-specific code
   ├── 📂 app                 # App config
   │   └── config.go          
   │
   ├── 📂 entity              # Core business models
   │   └── user.go            # User struct and validation logic
   │
   ├── 📂 usecase             # Business logic
   │   └── user_usecase.go    # CRUD operations for user
   │
   ├── 📂 repository          # Database interactions
   │   └── user_repo.go       # User data operations (Create, Read, Update, Delete)
   │
   ├── 📂 handler             # HTTP handlers (Gin controllers)
   │   └── user_handler.go    # Routes and request handling for user
   │
   └── 📂 router              # Router setup
        └── router.go         # Gin routes and setup            
```

---

### 🧠 How the Layers Interact
The typical flow follows:
```
handler -> usecase -> repository
```
**Why does `usecase -> repository` happen?**

In Clean Architecture, the **use case depends on an abstraction (interface)**, not the concrete implementation of the repository. This keeps the dependency direction correct:
- The **use case defines an interface** like `UserRepository`, describing required operations.
- The **repository layer implements this interface**, handling actual database operations.
- The **use case knows only the interface**, not the implementation — maintaining decoupling.

If the **use case directly called the repository struct**, that would break the dependency rule because a higher-level policy would depend on a lower-level detail.

So, there’s **no conceptual mistake** here — just ensure:
- The **repository interface is defined in the use case layer**.
- The **repository implementation stays in the infrastructure layer**.

This keeps everything flexible, scalable, and testable! 🚀

