import { UsersAPI } from "@repo/api-contract/client";

/**
 * API client instance for the Next.js app
 * Uses environment variable for API URL
 */
export const api = {
  users: new UsersAPI(
    process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1"
  ),
};
