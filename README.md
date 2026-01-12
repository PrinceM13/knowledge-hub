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

### Database

- **PostgreSQL** (Docker-based, local development)

Postgres is the only required data store at the moment.

### Planned / Future

- **Redis** – caching, queues, or performance experiments  
  _(not set up yet)_

---

## Local Development

### Prerequisites

- Docker & Docker Compose
- Node.js / Go (depending on the app you are running)

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

### Start Services

```bash
docker compose up -d
```
