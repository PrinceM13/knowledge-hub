# Local Development Guide

This document explains how to run the Knowledge Hub project locally and avoid common setup issues.

---

## Overview

For local development:

- Databases run inside Docker
- Applications (backend / frontend) may run either:
  - on the host machine (early stage), or
  - inside Docker (later stage)

Understanding **ports** and **where your app runs** is essential.

---

## PostgreSQL

### How PostgreSQL Is Set Up

- PostgreSQL runs **inside a Docker container**
- Inside Docker, Postgres listens on port **5432**
- Docker exposes it to your local machine on port **5433**

```text
Host Machine (Mac)
localhost:5433  ──▶  Docker Container:5432 (Postgres)
```

---

### Connecting From Your Local Machine

Use these settings when connecting from:
- TablePlus
- `psql` on your Mac
- Backend services running **outside Docker**

```text
Host: localhost
Port: 5433
Database: knowledge_hub
```

---

### Connecting From Inside Docker

If a service runs **inside Docker** (e.g. API container):

```text
Host: postgres
Port: 5432
Database: knowledge_hub
```

Docker containers communicate using **service names**, not `localhost`.

---

## Common Issues

### "role does not exist"

If Postgres works via `docker exec` but not from TablePlus or `psql -h`:

- You are likely connecting to a **different Postgres instance** on your host
- Check which process owns port 5432:

```bash
lsof -i :5432
```

If a local Postgres is running, always use **5433** for Docker Postgres.

---

## Useful Commands

Start services:
```bash
docker compose up -d
```

Stop services:
```bash
docker compose down
```

Connect to Postgres (host):
```bash
psql -h localhost -p 5433 -U kh_user knowledge_hub
```

Connect to Postgres (inside container):
```bash
docker exec -it knowledge_hub_postgres psql -U kh_user -d knowledge_hub
```

---

## Notes for Juniors

- Frontend apps **never connect directly to the database**
- Frontend → Backend API → Database
- Port confusion is normal — always ask *"Where is my app running?"*

This document will grow as more services are added.

