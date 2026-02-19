import type {
  CreateUserRequest,
  UserDetailResponse,
  UserListItem,
  ListUsersParams,
} from "../types/user";

/**
 * Error response structure from API
 */
interface ErrorResponse {
  message?: string;
  [key: string]: any;
}

/**
 * API Error class for handling API errors
 */
export class APIError extends Error {
  constructor(
    public status: number,
    message: string,
    public response?: any
  ) {
    super(message);
    this.name = "APIError";
  }
}

/**
 * Users API client
 * Matches endpoints in apps/api-go/internal/http/v1/user/handle.go
 */
export class UsersAPI {
  constructor(private baseURL: string) {}

  /**
   * Create a new user
   * POST /api/v1/users
   */
  async createUser(data: CreateUserRequest): Promise<UserDetailResponse> {
    const response = await fetch(`${this.baseURL}/users`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      const error = (await response.json().catch(() => ({}))) as ErrorResponse;
      throw new APIError(
        response.status,
        error.message || "Failed to create user",
        error
      );
    }

    return response.json() as Promise<UserDetailResponse>;
  }

  /**
   * List users with pagination
   * GET /api/v1/users?limit=20&offset=0
   */
  async listUsers(params?: ListUsersParams): Promise<UserListItem[]> {
    const searchParams = new URLSearchParams();

    if (params?.limit !== undefined) {
      searchParams.set("limit", params.limit.toString());
    }
    if (params?.offset !== undefined) {
      searchParams.set("offset", params.offset.toString());
    }

    const url = `${this.baseURL}/users${searchParams.toString() ? `?${searchParams}` : ""}`;

    const response = await fetch(url);

    if (!response.ok) {
      const error = (await response.json().catch(() => ({}))) as ErrorResponse;
      throw new APIError(
        response.status,
        error.message || "Failed to list users",
        error
      );
    }

    return response.json() as Promise<UserListItem[]>;
  }

  /**
   * Get user by ID
   * GET /api/v1/users/:id
   */
  async getUserById(id: number): Promise<UserDetailResponse> {
    const response = await fetch(`${this.baseURL}/users/${id}`);

    if (!response.ok) {
      const error = (await response.json().catch(() => ({}))) as ErrorResponse;
      throw new APIError(
        response.status,
        error.message || "Failed to get user",
        error
      );
    }

    return response.json() as Promise<UserDetailResponse>;
  }
}
