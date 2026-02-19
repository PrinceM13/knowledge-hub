/**
 * User model matching Go's internal/user/model.go
 */
export interface User {
  id: number;
  email: string;
  name: string;
  createdAt: string; // ISO 8601 date string (RFC3339 from Go)
}

/**
 * User list item response (minimal user data for lists)
 * Matches Go's UserListItemResponse
 */
export interface UserListItem {
  email: string;
  name: string;
}

/**
 * Create user request body
 * Matches Go's CreateUserRequest
 */
export interface CreateUserRequest {
  email: string;
  name: string;
}

/**
 * User detail response (full user data)
 * Matches Go's UserDetailResponse
 */
export interface UserDetailResponse {
  id: number;
  email: string;
  name: string;
  createdAt: string; // ISO 8601 date string
}

/**
 * Query parameters for listing users
 */
export interface ListUsersParams {
  limit?: number;
  offset?: number;
}
