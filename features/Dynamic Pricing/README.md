# 🛣️ VOID Core Service and Dynamic Routing Implementation

This directory contains the core service bootstrapping and foundational components for the VOID application, which is a microservice-oriented backend that supports dynamic features like real-time routing.

## 📁 Architecture Overview

The system is built on a standard Go microservice pattern using the **Gin** framework.

| Component | Directory (Implied) | Role in Routing/System |
| :--- | :--- | :--- |
| **Main Bootstrap** | `main.go` | Initializes all core services (DB, Cache, WS Manager) and registers the HTTP API routes. |
| **API Endpoints** | `internal/api/` | Handles HTTP requests for booking, status updates, and dynamic route queries. |
| **Database** | `internal/db/` | Manages persistent data, including Driver/User location, status, and historical data necessary for routing. |
| **Cache** | `internal/cache/` | Used for fast lookups of frequently accessed data (e.g., live vehicle positions or pricing factors). |
| **WebSockets** | `internal/ws/` | Handles real-time communication for driver updates and route guidance. **Essential for dynamic routing.** |
| **Configuration** | `config/` | Loads settings for database connection, server port, and external services. |

---

## 🛠️ Core Implementation Details

### 1. Service Initialization (`main.go`)

The entry point orchestrates the startup of all essential components:

* **Configuration:** `config.LoadConfig()` loads settings from files or environment variables.
* **Database (DB):** `db.InitDB(cfg)` sets up the GORM connection (defaulting to SQLite or PostgreSQL via `DATABASE_URL` environment variable).
    * `db.AutoMigrate(gormDB)` ensures the database schema (models like `User`, `Vehicle`) is up-to-date.
* **Cache:** `cache.NewCache(cfg)` initializes the caching layer (e.g., Redis) for quick data access.
* **Real-time Layer:** `ws.NewManager()` initializes and starts the WebSocket connection manager in a separate goroutine (`go manager.Run()`), allowing for low-latency communication crucial for updating routes dynamically.
* **HTTP Server:** `gin.Default()` creates the Gin router, and `api.RegisterRoutes(...)` maps all API endpoints, providing the interface for all routing logic.

### 2. API & Service Layer Interaction (Implied)

The actual dynamic routing logic (e.g., calculating the shortest or fastest route based on current traffic/demand) resides within the **`internal/api`** and **`internal/services`** (implied) packages.

* The API endpoints (`api.RegisterRoutes`) will expose methods like `/route/calculate` or `/trip/request`.
* These handlers call into a **Routing Service** which would:
    * Query the **DB** for driver and destination coordinates.
    * Query the **Cache** for real-time congestion data or demand factors.
    * Use a dedicated algorithm (potentially within a separate package like `internal/routing` or `features/dynamic_pricing`) to compute the optimal route.
    * The result (new route, estimated time) can be pushed to the client via the **WebSocket Manager**.

### 3. Scripting and Seeding (`seed_data.go`)

The `seed_data.go` file demonstrates the initial setup required for the routing system to function:

* It creates initial data for two sample **Drivers** (`Alice Driver`, `Bob Driver`) as `models.User` records with the role `"driver"`.
* It populates initial vehicle locations using the `models.Vehicle` model.
* **Key Insight:** The `Latitude` and `Longitude` fields in these models are the starting points for all dynamic routing calculations.

### 4. Simulation (`simulate_traffic.go` - Placeholder)

The empty `simulate_traffic.go` file is a clear placeholder for future or existing scripts responsible for generating or simulating real-time events, such as:

* Injecting **fake traffic or congestion data** into the Cache or DB.
* Simulating driver location updates, which are essential inputs for the dynamic routing engine.

This script demonstrates that the system is designed to consume external or simulated data streams to feed its dynamic capabilities.

---

Would you like to see a conceptual structure for the `internal/api/routes.go` file, showing how a dynamic routing endpoint might be registered?
