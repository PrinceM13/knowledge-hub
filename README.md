# Knowledge Hub

A monorepo for a personal knowledge management system.

This project is designed to:

- Learn and compare multiple frontend frameworks
- Practice real-world backend API design
- Serve as a training platform for juniors and stack-changers

---

## Repository Structure

- `apps/`  
  Runnable applications (API, frontend apps)

- `docker/`  
  Local development infrastructure (e.g. databases)

- `docs/`  
  Architecture notes, decisions, and training materials

---

## Tech Stack (Current)

### Backend

- **Go** (API server with clean architecture)
- **PostgreSQL** (Docker-based, local development)

### Frontend (Planned)

Starting with **Next.js**, then experimenting with:

- **Vue.js** – alternative web framework
- **Angular** – enterprise-grade framework
- **React Native** – mobile app for iOS/Android

The same backend API will serve all frontend implementations.

### Planned / Future

- **Redis** – caching, queues, or performance experiments  
  _(not set up yet)_

---

## Local Development

### Prerequisites

- Docker & Docker Compose
- **Go** (1.21+) for backend development
- **Node.js** (18+) for frontend development
- **pnpm** (8+) for Node.js package management

---

### PostgreSQL

PostgreSQL runs inside Docker and is exposed to the host machine.

**Connection info (from your local machine):**

- Host: `localhost`
- Port: `5433`
- Database: `knowledge_hub`

> Usernames and passwords are defined in the Docker configuration.

#### Port Explanation

PostgreSQL listens on port `5432` **inside the Docker container**.

Docker exposes it to the host machine on port `5433`:

- Use **port `5433`** when connecting from:
  - TablePlus
  - `psql` on your local machine
  - Backend services running outside Docker

- Use **port `5432`** only for services running **inside Docker**.

---

### Initial Setup

#### 1. Install Dependencies

Install all workspace dependencies:

```bash
pnpm install
```

This installs dependencies for all apps in the monorepo.

#### 2. Start Docker Services

Start PostgreSQL:

```bash
make db-up
# or
docker compose -f docker/docker-compose.yml up -d
```

#### 3. Run Next.js Development Server

```bash
make dev
# or
pnpm --filter web-next dev
```

The Next.js app will be available at `http://localhost:3000`

### Common Commands

#### Monorepo Management

```bash
# Install all dependencies
make install
# or
pnpm install

# Clean all node_modules and build artifacts
make clean
```

#### Next.js Development

```bash
# Start dev server (http://localhost:3000)
make dev

# Build for production
make build-next

# Run linter
make lint-next
```

#### Database Management

```bash
# Start PostgreSQL
make db-up

# Stop PostgreSQL
make db-down

# View logs
make db-logs

# Connect to database
make db-psql
```

#### Migrations

```bash
# Run migrations
make migrate-up

# Rollback last migration
make migrate-down

# Create new migration
make migrate-create name=migration_name
```

---

### Project Structure

```
knowledge-hub/
├── apps/
│   ├── api-go/                 # Go API server (Gin + Clean Architecture)
│   │   ├── cmd/
│   │   │   └── server/         # Main entry point
│   │   ├── internal/           # Business logic, handlers, repos
│   │   │   ├── app/            # Application layer
│   │   │   ├── config/         # Configuration
│   │   │   ├── db/             # Database repositories
│   │   │   ├── errors/         # Error handling
│   │   │   ├── health/         # Health check endpoints
│   │   │   ├── http/           # HTTP handlers & routing
│   │   │   ├── middleware/     # HTTP middleware
│   │   │   ├── server/         # HTTP server setup
│   │   │   ├── user/           # User domain logic
│   │   │   └── testutil/       # Test utilities
│   │   ├── migrations/         # SQL migrations
│   │   ├── go.mod
│   │   └── go.sum
│   │
│   ├── web-next/               # Next.js app (TypeScript + Tailwind CSS)
│   │   ├── src/                # Source code
│   │   │   ├── app/            # App Router pages
│   │   │   └── components/     # React components
│   │   ├── public/             # Static assets
│   │   ├── .env.example        # Environment variables template
│   │   ├── .env.local          # Local environment variables (gitignored)
│   │   ├── next.config.ts      # Next.js configuration
│   │   ├── tailwind.config.ts  # Tailwind configuration
│   │   ├── tsconfig.json       # TypeScript configuration
│   │   └── package.json        # Dependencies
│   │
│   ├── web-vue/                # Vue 3 (planned)
│   ├── web-svelte/             # SvelteKit (planned)
│   ├── web-angular/            # Angular (planned)
│   └── mobile-rn/              # React Native (planned)
│
├── packages/                   # Shared packages (future)
│   ├── api-contract/           # OpenAPI spec / shared types
│   └── ui-guidelines/          # Shared design system
│
├── docs/
│   └── local_dev.md            # Local development guide
│
├── docker/
│   └── docker-compose.yml      # PostgreSQL container
│
├── Makefile                    # Development commands
├── package.json                # Root workspace configuration
├── pnpm-workspace.yaml         # pnpm monorepo setup
├── .npmrc                      # pnpm settings
├── .gitignore
└── README.md
```
