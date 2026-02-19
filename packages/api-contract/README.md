# API Contract Package

TypeScript types and API client for the Knowledge Hub API.

## Usage

### Types Only

```typescript
import type { User, CreateUserRequest } from "@repo/api-contract/types";
```

### API Client

```typescript
import { UsersAPI } from "@repo/api-contract/client";

const api = new UsersAPI("http://localhost:8080/api/v1");

// Create user
const user = await api.createUser({
  email: "test@example.com",
  name: "Test User",
});

// List users
const users = await api.listUsers({ limit: 10, offset: 0 });

// Get user by ID
const user = await api.getUserById(1);
```

## API Endpoints

This package matches the Go API endpoints:

- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID

## Type Mapping

| Go Type     | TypeScript Type     |
| ----------- | ------------------- |
| `int64`     | `number`            |
| `string`    | `string`            |
| `time.Time` | `string` (ISO 8601) |
